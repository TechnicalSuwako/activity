// Code generated by astool. DO NOT EDIT.

package streams

import (
	typeemoji "github.com/go-fed/activity/streams/impl/toot/type_emoji"
	typehashtag "github.com/go-fed/activity/streams/impl/toot/type_hashtag"
	typeidentityproof "github.com/go-fed/activity/streams/impl/toot/type_identityproof"
	vocab "github.com/go-fed/activity/streams/vocab"
)

// IsOrExtendsTootEmoji returns true if the other provided type is the Emoji type
// or extends from the Emoji type.
func IsOrExtendsTootEmoji(other vocab.Type) bool {
	return typeemoji.IsOrExtendsEmoji(other)
}

// IsOrExtendsTootHashtag returns true if the other provided type is the Hashtag
// type or extends from the Hashtag type.
func IsOrExtendsTootHashtag(other vocab.Type) bool {
	return typehashtag.IsOrExtendsHashtag(other)
}

// IsOrExtendsTootIdentityProof returns true if the other provided type is the
// IdentityProof type or extends from the IdentityProof type.
func IsOrExtendsTootIdentityProof(other vocab.Type) bool {
	return typeidentityproof.IsOrExtendsIdentityProof(other)
}
