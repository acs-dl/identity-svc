/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type Positions struct {
	Key
	Attributes PositionsAttributes `json:"attributes"`
}
type PositionsResponse struct {
	Data     Positions `json:"data"`
	Included Included  `json:"included"`
}

type PositionsListResponse struct {
	Data     []Positions `json:"data"`
	Included Included    `json:"included"`
	Links    *Links      `json:"links"`
}

// MustPositions - returns Positions from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustPositions(key Key) *Positions {
	var positions Positions
	if c.tryFindEntry(key, &positions) {
		return &positions
	}
	return nil
}
