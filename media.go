// Package media is the Media content pack for togo: YouTubeVideo resource (views/likes/duration) — pairs with the compose queue for AI-drafted posts.
//
// A content pack bundles resource definitions (model + sqlc + Atlas + REST +
// GraphQL + a page) that a togo app pulls in. Generate the resource with
// `togo make:resource` in the host app; this plugin wires the provider hook.
package media

import "github.com/togo-framework/togo"

func init() {
	togo.RegisterProviderFunc("media", togo.PriorityService, func(k *togo.Kernel) error {
		return nil
	})
}
