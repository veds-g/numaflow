#!/usr/bin/env python3

#
# Filter swagger file
#

import os
from os import path
import sys
import json


def main():
    if sys.stdin.isatty():
        print("ERROR: swagger json needs to be piped in as stdin")
        exit(1)
    if len(sys.argv) < 2:
        print("ERROR: definition prefix needs to be provided")
        exit(1)
    prefix = sys.argv[1]

    try:
        swagger = json.load(sys.stdin)
    except Exception as e:
        print("ERROR: not a valid json input - {0}".format(e))
        exit(1)

    defs = swagger["definitions"]
    for k in list(defs.keys()):
        if not k.startswith(prefix):
            del defs[k]
            continue

        if k in [
            "io.numaproj.numaflow.v1alpha1.Blackhole",
            "io.numaproj.numaflow.v1alpha1.Log",
            "io.numaproj.numaflow.v1alpha1.NoStore",
            "io.numaproj.numaflow.v1alpha1.ServeSink",
            "io.numaproj.numaflow.v1alpha1.ServingSource",
        ]:
            defs[k]["allOf"] = []

    json_object = json.dumps(swagger, indent=4)
    print(json_object)


if __name__ == "__main__":
    main()
