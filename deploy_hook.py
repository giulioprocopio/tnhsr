from itertools import chain
import os
from pathlib import Path

DIR = Path(__file__).parent.absolute()


def read_env_file(path):
    with open(path, 'r') as f:
        for line in f.readlines():
            # Comment or empty line.
            sline = line.strip()
            if sline.startswith('#') or not sline:
                continue

            # Key-value pair.
            key, value = sline.split('=', 1)
            while len(value) > 1 and value[0] == value[-1] and value[0] in [
                    '"', "'"
            ]:
                value = value[1:-1]

            yield key, value


def main():
    errno = 0

    env_sample = dict(read_env_file(DIR / '.env.sample'))
    print('Must set: ', end='')
    print(', '.join(env_sample.keys()))

    env = {}
    for path in chain(DIR.glob('*.env'), DIR.glob('.env.*')):
        if path.name == '.env.sample':
            continue
        env.update(read_env_file(path))

    env.update(os.environ)

    for key, _ in env_sample.items():
        if key not in env:
            print(f'Error: `{key}` is not set.')
            errno = 1
    else:
        if errno == 0:
            print('All env vars are set.')

    return errno


if __name__ == '__main__':
    exit(main())
