Mozsvc_util
===========

A few go utility functions I found useful

util.MZConfig - Reads a ini like file of settings.

util.GenUUID4() - simple UUID4 generator / parser

util.GetAWSPublicHostname() - Returns the AWS Public Hostname

util.HekaLogger - Interface to local heka logging writer.

util.Metrics - simple statsd like metric recorder functions.

These are pretty much for my own use and abuse at this time.
As I get time/inclination/wild hair, I'll improve them and maybe make
them more generally useful. Heck, maybe I'll even add proper unit and
interface tests to them. As it stands, I tend to modify these when I
actually use them in projects, so they're hardly stand-alone.

The dependencies are available in go-prod.deps.
