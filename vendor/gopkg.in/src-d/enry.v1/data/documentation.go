package data

// CODE GENERATED AUTOMATICALLY WITH gopkg.in/src-d/enry.v1/internal/code-generator
// THIS FILE SHOULD NOT BE EDITED BY HAND
// Extracted from github/linguist commit: 4cd558c37482e8d2c535d8107f2d11b49afbc5b5

import "gopkg.in/toqueteos/substring.v1"

var DocumentationMatchers = substring.Or(
	substring.Regexp(`^[Dd]ocs?/`),
	substring.Regexp(`(^|/)[Dd]ocumentation/`),
	substring.Regexp(`(^|/)[Jj]avadoc/`),
	substring.Regexp(`^[Mm]an/`),
	substring.Regexp(`^[Ee]xamples/`),
	substring.Regexp(`^[Dd]emos?/`),
	substring.Regexp(`(^|/)CHANGE(S|LOG)?(\.|$)`),
	substring.Regexp(`(^|/)CONTRIBUTING(\.|$)`),
	substring.Regexp(`(^|/)COPYING(\.|$)`),
	substring.Regexp(`(^|/)INSTALL(\.|$)`),
	substring.Regexp(`(^|/)LICEN[CS]E(\.|$)`),
	substring.Regexp(`(^|/)[Ll]icen[cs]e(\.|$)`),
	substring.Regexp(`(^|/)README(\.|$)`),
	substring.Regexp(`(^|/)[Rr]eadme(\.|$)`),
	substring.Regexp(`^[Ss]amples?/`),
)
