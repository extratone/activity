package vocab

import "net/url"

// ActivityStreamsIconPropertyIterator represents a single value for the "icon"
// property.
type ActivityStreamsIconPropertyIterator interface {
	// GetActivityStreamsImage returns the value of this property. When
	// IsActivityStreamsImage returns false, GetActivityStreamsImage will
	// return an arbitrary value.
	GetActivityStreamsImage() ActivityStreamsImage
	// GetActivityStreamsLink returns the value of this property. When
	// IsActivityStreamsLink returns false, GetActivityStreamsLink will
	// return an arbitrary value.
	GetActivityStreamsLink() ActivityStreamsLink
	// GetActivityStreamsMention returns the value of this property. When
	// IsActivityStreamsMention returns false, GetActivityStreamsMention
	// will return an arbitrary value.
	GetActivityStreamsMention() ActivityStreamsMention
	// GetIRI returns the IRI of this property. When IsIRI returns false,
	// GetIRI will return an arbitrary value.
	GetIRI() *url.URL
	// GetType returns the value in this property as a Type. Returns nil if
	// the value is not an ActivityStreams type, such as an IRI or another
	// value.
	GetType() Type
	// HasAny returns true if any of the different values is set.
	HasAny() bool
	// IsActivityStreamsImage returns true if this property has a type of
	// "Image". When true, use the GetActivityStreamsImage and
	// SetActivityStreamsImage methods to access and set this property.
	IsActivityStreamsImage() bool
	// IsActivityStreamsLink returns true if this property has a type of
	// "Link". When true, use the GetActivityStreamsLink and
	// SetActivityStreamsLink methods to access and set this property.
	IsActivityStreamsLink() bool
	// IsActivityStreamsMention returns true if this property has a type of
	// "Mention". When true, use the GetActivityStreamsMention and
	// SetActivityStreamsMention methods to access and set this property.
	IsActivityStreamsMention() bool
	// IsIRI returns true if this property is an IRI. When true, use GetIRI
	// and SetIRI to access and set this property
	IsIRI() bool
	// JSONLDContext returns the JSONLD URIs required in the context string
	// for this property and the specific values that are set. The value
	// in the map is the alias used to import the property's value or
	// values.
	JSONLDContext() map[string]string
	// KindIndex computes an arbitrary value for indexing this kind of value.
	// This is a leaky API detail only for folks looking to replace the
	// go-fed implementation. Applications should not use this method.
	KindIndex() int
	// LessThan compares two instances of this property with an arbitrary but
	// stable comparison. Applications should not use this because it is
	// only meant to help alternative implementations to go-fed to be able
	// to normalize nonfunctional properties.
	LessThan(o ActivityStreamsIconPropertyIterator) bool
	// Name returns the name of this property: "ActivityStreamsIcon".
	Name() string
	// Next returns the next iterator, or nil if there is no next iterator.
	Next() ActivityStreamsIconPropertyIterator
	// Prev returns the previous iterator, or nil if there is no previous
	// iterator.
	Prev() ActivityStreamsIconPropertyIterator
	// SetActivityStreamsImage sets the value of this property. Calling
	// IsActivityStreamsImage afterwards returns true.
	SetActivityStreamsImage(v ActivityStreamsImage)
	// SetActivityStreamsLink sets the value of this property. Calling
	// IsActivityStreamsLink afterwards returns true.
	SetActivityStreamsLink(v ActivityStreamsLink)
	// SetActivityStreamsMention sets the value of this property. Calling
	// IsActivityStreamsMention afterwards returns true.
	SetActivityStreamsMention(v ActivityStreamsMention)
	// SetIRI sets the value of this property. Calling IsIRI afterwards
	// returns true.
	SetIRI(v *url.URL)
	// SetType attempts to set the property for the arbitrary type. Returns an
	// error if it is not a valid type to set on this property.
	SetType(t Type) error
}

// Indicates an entity that describes an icon for this object. The image should
// have an aspect ratio of one (horizontal) to one (vertical) and should be
// suitable for presentation at a small size.
//
// Example 79 (https://www.w3.org/TR/activitystreams-vocabulary/#ex77-jsonld):
//   {
//     "content": "This is all there is.",
//     "icon": {
//       "height": 16,
//       "name": "Note icon",
//       "type": "Image",
//       "url": "http://example.org/note.png",
//       "width": 16
//     },
//     "summary": "A simple note",
//     "type": "Note"
//   }
//
// Example 80 (https://www.w3.org/TR/activitystreams-vocabulary/#ex78-jsonld):
//   {
//     "content": "A simple note",
//     "icon": [
//       {
//         "height": 16,
//         "summary": "Note (16x16)",
//         "type": "Image",
//         "url": "http://example.org/note1.png",
//         "width": 16
//       },
//       {
//         "height": 32,
//         "summary": "Note (32x32)",
//         "type": "Image",
//         "url": "http://example.org/note2.png",
//         "width": 32
//       }
//     ],
//     "summary": "A simple note",
//     "type": "Note"
//   }
type ActivityStreamsIconProperty interface {
	// AppendActivityStreamsImage appends a Image value to the back of a list
	// of the property "icon". Invalidates iterators that are traversing
	// using Prev.
	AppendActivityStreamsImage(v ActivityStreamsImage)
	// AppendActivityStreamsLink appends a Link value to the back of a list of
	// the property "icon". Invalidates iterators that are traversing
	// using Prev.
	AppendActivityStreamsLink(v ActivityStreamsLink)
	// AppendActivityStreamsMention appends a Mention value to the back of a
	// list of the property "icon". Invalidates iterators that are
	// traversing using Prev.
	AppendActivityStreamsMention(v ActivityStreamsMention)
	// AppendIRI appends an IRI value to the back of a list of the property
	// "icon"
	AppendIRI(v *url.URL)
	// PrependType prepends an arbitrary type value to the front of a list of
	// the property "icon". Invalidates iterators that are traversing
	// using Prev. Returns an error if the type is not a valid one to set
	// for this property.
	AppendType(t Type) error
	// At returns the property value for the specified index. Panics if the
	// index is out of bounds.
	At(index int) ActivityStreamsIconPropertyIterator
	// Begin returns the first iterator, or nil if empty. Can be used with the
	// iterator's Next method and this property's End method to iterate
	// from front to back through all values.
	Begin() ActivityStreamsIconPropertyIterator
	// Empty returns returns true if there are no elements.
	Empty() bool
	// End returns beyond-the-last iterator, which is nil. Can be used with
	// the iterator's Next method and this property's Begin method to
	// iterate from front to back through all values.
	End() ActivityStreamsIconPropertyIterator
	// JSONLDContext returns the JSONLD URIs required in the context string
	// for this property and the specific values that are set. The value
	// in the map is the alias used to import the property's value or
	// values.
	JSONLDContext() map[string]string
	// KindIndex computes an arbitrary value for indexing this kind of value.
	// This is a leaky API method specifically needed only for alternate
	// implementations for go-fed. Applications should not use this
	// method. Panics if the index is out of bounds.
	KindIndex(idx int) int
	// Len returns the number of values that exist for the "icon" property.
	Len() (length int)
	// Less computes whether another property is less than this one. Mixing
	// types results in a consistent but arbitrary ordering
	Less(i, j int) bool
	// LessThan compares two instances of this property with an arbitrary but
	// stable comparison. Applications should not use this because it is
	// only meant to help alternative implementations to go-fed to be able
	// to normalize nonfunctional properties.
	LessThan(o ActivityStreamsIconProperty) bool
	// Name returns the name of this property: "icon".
	Name() string
	// PrependActivityStreamsImage prepends a Image value to the front of a
	// list of the property "icon". Invalidates all iterators.
	PrependActivityStreamsImage(v ActivityStreamsImage)
	// PrependActivityStreamsLink prepends a Link value to the front of a list
	// of the property "icon". Invalidates all iterators.
	PrependActivityStreamsLink(v ActivityStreamsLink)
	// PrependActivityStreamsMention prepends a Mention value to the front of
	// a list of the property "icon". Invalidates all iterators.
	PrependActivityStreamsMention(v ActivityStreamsMention)
	// PrependIRI prepends an IRI value to the front of a list of the property
	// "icon".
	PrependIRI(v *url.URL)
	// PrependType prepends an arbitrary type value to the front of a list of
	// the property "icon". Invalidates all iterators. Returns an error if
	// the type is not a valid one to set for this property.
	PrependType(t Type) error
	// Remove deletes an element at the specified index from a list of the
	// property "icon", regardless of its type. Panics if the index is out
	// of bounds. Invalidates all iterators.
	Remove(idx int)
	// Serialize converts this into an interface representation suitable for
	// marshalling into a text or binary format. Applications should not
	// need this function as most typical use cases serialize types
	// instead of individual properties. It is exposed for alternatives to
	// go-fed implementations to use.
	Serialize() (interface{}, error)
	// SetActivityStreamsImage sets a Image value to be at the specified index
	// for the property "icon". Panics if the index is out of bounds.
	// Invalidates all iterators.
	SetActivityStreamsImage(idx int, v ActivityStreamsImage)
	// SetActivityStreamsLink sets a Link value to be at the specified index
	// for the property "icon". Panics if the index is out of bounds.
	// Invalidates all iterators.
	SetActivityStreamsLink(idx int, v ActivityStreamsLink)
	// SetActivityStreamsMention sets a Mention value to be at the specified
	// index for the property "icon". Panics if the index is out of
	// bounds. Invalidates all iterators.
	SetActivityStreamsMention(idx int, v ActivityStreamsMention)
	// SetIRI sets an IRI value to be at the specified index for the property
	// "icon". Panics if the index is out of bounds.
	SetIRI(idx int, v *url.URL)
	// Swap swaps the location of values at two indices for the "icon"
	// property.
	Swap(i, j int)
}