from itertools import chain
import os
from pathlib import Path

DIR = Path(__file__).parent.absolute()
ENV_SAMPLE = DIR / '.env.sample'


def iter_env_file(path):
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


def iter_env_files_paths():
    for path in chain(DIR.glob('*.env'), DIR.glob('.env.*')):
        if path == ENV_SAMPLE:
            continue

        yield path


def main():
    errno = 0

    assert ENV_SAMPLE.exists()
    env_sample = dict(iter_env_file(ENV_SAMPLE))
    print('must set: ', end='')
    print(', '.join(env_sample.keys()))

    env = {}
    for path in iter_env_files_paths():
        print(f'reading {path.relative_to(DIR)}')
        env.update(iter_env_file(path))

    print('reading current env')
    env.update(os.environ)

    for key, _ in env_sample.items():
        if key not in env:
            print(f'error: `{key}` is not set')
            errno = 1
    else:
        if errno == 0:
            print('all env vars are set')

    return errno


if __name__ == '__main__':
    exit(main())
