/*
Command brigade is a toolkit to list and sync S3 buckets.

It can create a listing of all the keys in a bucket, compute the difference
between two listing, slice a listing into sub-listings and synchronize
a listing between a source bucket onto a destination bucket, using
PutCopy.

It also provides a convenience `backup` command that performs all those steps
automatically, using a third S3 bucket to keep state between executions. This
command is most appropriate for periodic backup jobs of an S3 bucket to another.

The motivation behind command brigade is to keep a copy of an S3 bucket
accessible from an incompatible set of credentials for the original bucket.
In a scenario where the original bucket is compromised and destroyed, the
copy would be up and relatively fresh, while inaccessible by an attacker.

    list     Lists the keys in an S3 bucket.
    sync     Syncs the keys from a source S3 bucket to another.
    slice    Slice an S3 key listing into multiple sub-listings.
    diff     Generates a differential listing of S3 keys.
    backup   Executes list, diff and sync from a source to a destination bucket.
    help, h  Shows a list of commands or help for one command

*/
package main
