import json, re, sys

BASE = '/Users/papapin777/Documents/courses/robotics/'
LESSONS = json.load(open(BASE + 'course.json'))['lessons']
EL = re.compile(r'^:::(\w+)')
COMMENT = re.compile(r'^(//|#)\s*([\w./-]+\.\w+)\s*(.*)$')
CODE_LANGS = ('cpp', 'text', 'c', 'ino')


def parse_items(lines):
    items, buf, i = [], [], 0
    in_fence = False
    while i < len(lines):
        ln = lines[i]
        if ln.startswith('```'):
            in_fence = not in_fence
        m = EL.match(ln)
        if m and not in_fence:
            if buf:
                items.append(('text', None, buf))
                buf = []
            j = i
            el = [ln]
            j += 1
            while j < len(lines) and lines[j] != ':::':
                el.append(lines[j])
                j += 1
            el.append(':::')
            items.append(('el', m.group(1), el))
            i = j + 1
            continue
        buf.append(ln)
        i += 1
    if buf:
        items.append(('text', None, buf))
    return items


def render(items):
    out = []
    for kind, name, ls in items:
        text = '\n'.join(ls).strip('\n')
        if text.strip():
            out.append(text)
    return '\n\n'.join(out)


def fix_fences(md, lesson):
    lines = md.split('\n')
    out = []
    i = 0
    while i < len(lines):
        ln = lines[i]
        out.append(ln)
        m = re.match(r'^```(\w[\w-]*)$', ln)
        if m and m.group(1) in CODE_LANGS and i + 1 < len(lines):
            first = lines[i + 1]
            cm = COMMENT.match(first)
            if cm:
                out.append(cm.group(1) + ' ' + cm.group(2))
                if cm.group(3):
                    out.append(cm.group(1) + ' ' + cm.group(3).strip('() '))
                i += 2
                continue
            else:
                print('NO FILE COMMENT', lesson, 'near line', i + 1, repr(first))
        # skip over fence body so closing fence is not treated as an opener
        if re.match(r'^```\w', ln):
            i += 1
            while i < len(lines) and not lines[i].startswith('```'):
                if out and i > 0:
                    pass
                out.append(lines[i])
                i += 1
            if i < len(lines):
                out.append(lines[i])
            i += 1
            continue
        i += 1
    return '\n'.join(out)


def relocate(path, lesson):
    md = open(path).read()
    head, *rest = re.split(r'(?m)^## ', md)
    names, bodies = [], {}
    for r in rest:
        name, _, body = r.partition('\n')
        names.append(name.strip())
        bodies[name.strip()] = parse_items(body.split('\n'))
    dest = {'hook': "Why This Matters", 'mistake': 'Common Mistakes', 'challenge': 'Try It Yourself',
            'homework': "What's Next", 'bridge': "What's Next"}
    moved = {k: [] for k in dest}
    for n in names:
        keep = []
        for it in bodies[n]:
            if it[0] == 'el' and it[1] in dest and dest[it[1]] != n:
                moved[it[1]].append(it)
            else:
                keep.append(it)
        bodies[n] = keep
    bodies[dest['hook']] = moved['hook'] + bodies[dest['hook']]
    bodies['Common Mistakes'] = bodies['Common Mistakes'] + moved['mistake']
    bodies['Try It Yourself'] = bodies['Try It Yourself'] + moved['challenge']
    bodies["What's Next"] = moved['bridge'] + bodies["What's Next"] + moved['homework']
    out = head.rstrip('\n') + '\n\n'
    for n in names:
        out += '## ' + n + '\n\n' + render(bodies[n]) + '\n\n'
    out = out.rstrip('\n') + '\n'
    out = fix_fences(out, lesson)
    open(path, 'w').write(out)


def report(path, lesson):
    md = open(path).read()
    lj = json.load(open(BASE + 'lessons/' + lesson + '/lesson.json'))
    for c in lj['concepts']:
        cid = c['id']
        miss = []
        for el in ('example', 'mistake', 'checkpoint', 'analogy'):
            if not re.search(r':::' + el + r' \{[^}]*concept="' + cid + '"', md):
                miss.append(el)
        if miss:
            print('MISSING', lesson, cid, miss)
    for h in ('What You\'ll Learn', 'Why This Matters', 'Walkthrough', 'Common Mistakes', 'Try It Yourself',
              'Key Takeaways', "What's Next", 'Further Reading'):
        if '\n## ' + h + '\n' not in md:
            print('NO HEADING', lesson, h)


for l in LESSONS:
    p = BASE + 'lessons/' + l + '/lesson.md'
    if len(sys.argv) > 1 and sys.argv[1] == 'fix':
        relocate(p, l)
    report(p, l)
