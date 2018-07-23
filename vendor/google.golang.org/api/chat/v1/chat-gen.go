// Package chat provides access to the Hangouts Chat API.
//
// See https://developers.google.com/hangouts/chat
//
// Usage example:
//
//   import "google.golang.org/api/chat/v1"
//   ...
//   chatService, err := chat.New(oauthHttpClient)
package chat // import "google.golang.org/api/chat/v1"

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	context "golang.org/x/net/context"
	ctxhttp "golang.org/x/net/context/ctxhttp"
	gensupport "google.golang.org/api/gensupport"
	googleapi "google.golang.org/api/googleapi"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Always reference these packages, just in case the auto-generated code
// below doesn't.
var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = gensupport.MarshalJSON
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace
var _ = context.Canceled
var _ = ctxhttp.Do

const apiId = "chat:v1"
const apiName = "chat"
const apiVersion = "v1"
const basePath = "https://chat.googleapis.com/"

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.Spaces = NewSpacesService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Spaces *SpacesService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewSpacesService(s *Service) *SpacesService {
	rs := &SpacesService{s: s}
	rs.Members = NewSpacesMembersService(s)
	rs.Messages = NewSpacesMessagesService(s)
	return rs
}

type SpacesService struct {
	s *Service

	Members *SpacesMembersService

	Messages *SpacesMessagesService
}

func NewSpacesMembersService(s *Service) *SpacesMembersService {
	rs := &SpacesMembersService{s: s}
	return rs
}

type SpacesMembersService struct {
	s *Service
}

func NewSpacesMessagesService(s *Service) *SpacesMessagesService {
	rs := &SpacesMessagesService{s: s}
	return rs
}

type SpacesMessagesService struct {
	s *Service
}

// ActionParameter: List of string parameters to supply when the action
// method is invoked.
// For example, consider three snooze buttons: snooze now, snooze 1
// day,
// snooze next week. You might use action method = snooze(), passing
// the
// snooze type and snooze time in the list of string parameters.
type ActionParameter struct {
	// Key: The name of the parameter for the action script.
	Key string `json:"key,omitempty"`

	// Value: The value of the parameter.
	Value string `json:"value,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Key") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Key") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ActionParameter) MarshalJSON() ([]byte, error) {
	type NoMethod ActionParameter
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ActionResponse: Parameters that a bot can use to configure how it's
// response is posted.
type ActionResponse struct {
	// Type: The type of bot response.
	//
	// Possible values:
	//   "TYPE_UNSPECIFIED" - Default type; will be handled as NEW_MESSAGE.
	//   "NEW_MESSAGE" - Post as a new message in the topic.
	//   "UPDATE_MESSAGE" - Update the bot's own message. (Only after
	// CARD_CLICKED events.)
	//   "REQUEST_CONFIG" - Privately ask the user for additional auth or
	// config.
	Type string `json:"type,omitempty"`

	// Url: URL for users to auth or config. (Only for REQUEST_CONFIG
	// response types.)
	Url string `json:"url,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Type") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Type") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ActionResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ActionResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Annotation: Annotations associated with the plain-text body of the
// message.
//
// Example plain-text message body:
// ```
// Hello @FooBot how are you!"
// ```
//
// The corresponding annotations metadata:
// ```
// "annotations":[{
//   "type":"USER_MENTION",
//   "startIndex":6,
//   "length":7,
//   "userMention": {
//     "user": {
//       "name":"users/107946847022116401880",
//       "displayName":"FooBot",
//       "avatarUrl":"https://goo.gl/aeDtrS",
//       "type":"BOT"
//     },
//     "type":"MENTION"
//    }
// }]
// ```
type Annotation struct {
	// Length: Length of the substring in the plain-text message body this
	// annotation
	// corresponds to.
	Length int64 `json:"length,omitempty"`

	// StartIndex: Start index (0-based, inclusive) in the plain-text
	// message body this
	// annotation corresponds to.
	StartIndex int64 `json:"startIndex,omitempty"`

	// Type: The type of this annotation.
	//
	// Possible values:
	//   "ANNOTATION_TYPE_UNSPECIFIED" - Default value for the enum. DO NOT
	// USE.
	//   "USER_MENTION" - A user is mentioned.
	Type string `json:"type,omitempty"`

	// UserMention: The metadata of user mention.
	UserMention *UserMentionMetadata `json:"userMention,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Length") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Length") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Annotation) MarshalJSON() ([]byte, error) {
	type NoMethod Annotation
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Button: A button. Can be a text button or an image button.
type Button struct {
	// ImageButton: A button with image and onclick action.
	ImageButton *ImageButton `json:"imageButton,omitempty"`

	// TextButton: A button with text and onclick action.
	TextButton *TextButton `json:"textButton,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ImageButton") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ImageButton") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Button) MarshalJSON() ([]byte, error) {
	type NoMethod Button
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Card: A card is a UI element that can contain UI widgets such as
// texts, images.
type Card struct {
	// CardActions: The actions of this card.
	CardActions []*CardAction `json:"cardActions,omitempty"`

	// Header: The header of the card. A header usually contains a title and
	// an image.
	Header *CardHeader `json:"header,omitempty"`

	// Name: Name of the card.
	Name string `json:"name,omitempty"`

	// Sections: Sections are separated by a line divider.
	Sections []*Section `json:"sections,omitempty"`

	// ForceSendFields is a list of field names (e.g. "CardActions") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CardActions") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Card) MarshalJSON() ([]byte, error) {
	type NoMethod Card
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// CardAction: A card action is
// the action associated with the card. For an invoice card, a
// typical action would be: delete invoice, email invoice or open
// the
// invoice in browser.
type CardAction struct {
	// ActionLabel: The label used to be displayed in the action menu item.
	ActionLabel string `json:"actionLabel,omitempty"`

	// OnClick: The onclick action for this action item.
	OnClick *OnClick `json:"onClick,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ActionLabel") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ActionLabel") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *CardAction) MarshalJSON() ([]byte, error) {
	type NoMethod CardAction
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type CardHeader struct {
	// ImageStyle: The image's type (e.g. square border or circular border).
	//
	// Possible values:
	//   "IMAGE_STYLE_UNSPECIFIED"
	//   "IMAGE" - Square border.
	//   "AVATAR" - Circular border.
	ImageStyle string `json:"imageStyle,omitempty"`

	// ImageUrl: The URL of the image in the card header.
	ImageUrl string `json:"imageUrl,omitempty"`

	// Subtitle: The subtitle of the card header.
	Subtitle string `json:"subtitle,omitempty"`

	// Title: The title must be specified. The header has a fixed height: if
	// both a
	// title and subtitle is specified, each will take up 1 line. If only
	// the
	// title is specified, it will take up both lines.
	Title string `json:"title,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ImageStyle") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ImageStyle") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *CardHeader) MarshalJSON() ([]byte, error) {
	type NoMethod CardHeader
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// DeprecatedEvent: Hangouts Chat events.
type DeprecatedEvent struct {
	// Action: The form action data associated with an interactive card that
	// was clicked.
	// Only populated for
	// CARD_CLICKED events.
	// See the [Interactive Cards
	// guide](/hangouts/chat/how-tos/cards-onclick) for
	// more information.
	Action *FormAction `json:"action,omitempty"`

	// ConfigCompleteRedirectUrl: The URL the bot should redirect the user
	// to after they have completed an
	// authorization or configuration flow outside of Hangouts Chat. See
	// the
	// [Authorizing access to 3p services
	// guide](/hangouts/chat/how-tos/auth-3p)
	// for more information.
	ConfigCompleteRedirectUrl string `json:"configCompleteRedirectUrl,omitempty"`

	// EventTime: The timestamp indicating when the event was dispatched.
	EventTime string `json:"eventTime,omitempty"`

	// Message: The message that triggered the event, if applicable.
	Message *Message `json:"message,omitempty"`

	// Space: The room or DM in which the event occurred.
	Space *Space `json:"space,omitempty"`

	// ThreadKey: The bot-defined key for the thread related to the event.
	// See the
	// thread_key field of the
	// `spaces.message.create` request for more information.
	ThreadKey string `json:"threadKey,omitempty"`

	// Token: A secret value that bots can use to verify if a request is
	// from Google. The
	// token is randomly generated by Google, remains static, and can be
	// obtained
	// from the Hangouts Chat API configuration page in the Cloud
	// Console.
	// Developers can revoke/regenerate it if needed from the same page.
	Token string `json:"token,omitempty"`

	// Type: The type of the event.
	//
	// Possible values:
	//   "UNSPECIFIED" - Default value for the enum. DO NOT USE.
	//   "MESSAGE" - A message was sent in a room or direct message.
	//   "ADDED_TO_SPACE" - The bot was added to a room or DM.
	//   "REMOVED_FROM_SPACE" - The bot was removed from a room or DM.
	//   "CARD_CLICKED" - The bot's interactive card was clicked.
	Type string `json:"type,omitempty"`

	// User: The user that triggered the event.
	User *User `json:"user,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Action") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Action") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *DeprecatedEvent) MarshalJSON() ([]byte, error) {
	type NoMethod DeprecatedEvent
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Empty: A generic empty message that you can re-use to avoid defining
// duplicated
// empty messages in your APIs. A typical example is to use it as the
// request
// or the response type of an API method. For instance:
//
//     service Foo {
//       rpc Bar(google.protobuf.Empty) returns
// (google.protobuf.Empty);
//     }
//
// The JSON representation for `Empty` is empty JSON object `{}`.
type Empty struct {
	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`
}

// FormAction: A form action describes the behavior when the form is
// submitted.
// For example, an Apps Script can be invoked to handle the form.
type FormAction struct {
	// ActionMethodName: Apps Script function to invoke when the containing
	// element is
	// clicked/activated.
	ActionMethodName string `json:"actionMethodName,omitempty"`

	// Parameters: List of action parameters.
	Parameters []*ActionParameter `json:"parameters,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ActionMethodName") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ActionMethodName") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *FormAction) MarshalJSON() ([]byte, error) {
	type NoMethod FormAction
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Image: An image that is specified by a URL and can have an onclick
// action.
type Image struct {
	// AspectRatio: The aspect ratio of this image (width/height).
	AspectRatio float64 `json:"aspectRatio,omitempty"`

	// ImageUrl: The URL of the image.
	ImageUrl string `json:"imageUrl,omitempty"`

	// OnClick: The onclick action.
	OnClick *OnClick `json:"onClick,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AspectRatio") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AspectRatio") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Image) MarshalJSON() ([]byte, error) {
	type NoMethod Image
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

func (s *Image) UnmarshalJSON(data []byte) error {
	type NoMethod Image
	var s1 struct {
		AspectRatio gensupport.JSONFloat64 `json:"aspectRatio"`
		*NoMethod
	}
	s1.NoMethod = (*NoMethod)(s)
	if err := json.Unmarshal(data, &s1); err != nil {
		return err
	}
	s.AspectRatio = float64(s1.AspectRatio)
	return nil
}

// ImageButton: An image button with an onclick action.
type ImageButton struct {
	// Icon: The icon specified by an enum that indices to an icon provided
	// by Chat
	// API.
	//
	// Possible values:
	//   "ICON_UNSPECIFIED"
	//   "AIRPLANE"
	//   "BOOKMARK"
	//   "BUS"
	//   "CAR"
	//   "CLOCK"
	//   "CONFIRMATION_NUMBER_ICON"
	//   "DOLLAR"
	//   "DESCRIPTION"
	//   "EMAIL"
	//   "EVENT_PERFORMER"
	//   "EVENT_SEAT"
	//   "FLIGHT_ARRIVAL"
	//   "FLIGHT_DEPARTURE"
	//   "HOTEL"
	//   "HOTEL_ROOM_TYPE"
	//   "INVITE"
	//   "MAP_PIN"
	//   "MEMBERSHIP"
	//   "MULTIPLE_PEOPLE"
	//   "OFFER"
	//   "PERSON"
	//   "PHONE"
	//   "RESTAURANT_ICON"
	//   "SHOPPING_CART"
	//   "STAR"
	//   "STORE"
	//   "TICKET"
	//   "TRAIN"
	//   "VIDEO_CAMERA"
	//   "VIDEO_PLAY"
	Icon string `json:"icon,omitempty"`

	// IconUrl: The icon specified by a URL.
	IconUrl string `json:"iconUrl,omitempty"`

	// Name: The name of this image_button which will be used for
	// accessibility.
	// Default value will be provided if developers don't specify.
	Name string `json:"name,omitempty"`

	// OnClick: The onclick action.
	OnClick *OnClick `json:"onClick,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Icon") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Icon") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ImageButton) MarshalJSON() ([]byte, error) {
	type NoMethod ImageButton
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// KeyValue: A UI element contains a key (label) and a value (content).
// And this
// element may also contain some actions such as onclick button.
type KeyValue struct {
	// BottomLabel: The text of the bottom label. Formatted text supported.
	BottomLabel string `json:"bottomLabel,omitempty"`

	// Button: A button that can be clicked to trigger an action.
	Button *Button `json:"button,omitempty"`

	// Content: The text of the content. Formatted text supported and always
	// required.
	Content string `json:"content,omitempty"`

	// ContentMultiline: If the content should be multiline.
	ContentMultiline bool `json:"contentMultiline,omitempty"`

	// Icon: An enum value that will be replaced by the Chat API with
	// the
	// corresponding icon image.
	//
	// Possible values:
	//   "ICON_UNSPECIFIED"
	//   "AIRPLANE"
	//   "BOOKMARK"
	//   "BUS"
	//   "CAR"
	//   "CLOCK"
	//   "CONFIRMATION_NUMBER_ICON"
	//   "DOLLAR"
	//   "DESCRIPTION"
	//   "EMAIL"
	//   "EVENT_PERFORMER"
	//   "EVENT_SEAT"
	//   "FLIGHT_ARRIVAL"
	//   "FLIGHT_DEPARTURE"
	//   "HOTEL"
	//   "HOTEL_ROOM_TYPE"
	//   "INVITE"
	//   "MAP_PIN"
	//   "MEMBERSHIP"
	//   "MULTIPLE_PEOPLE"
	//   "OFFER"
	//   "PERSON"
	//   "PHONE"
	//   "RESTAURANT_ICON"
	//   "SHOPPING_CART"
	//   "STAR"
	//   "STORE"
	//   "TICKET"
	//   "TRAIN"
	//   "VIDEO_CAMERA"
	//   "VIDEO_PLAY"
	Icon string `json:"icon,omitempty"`

	// IconUrl: The icon specified by a URL.
	IconUrl string `json:"iconUrl,omitempty"`

	// OnClick: The onclick action. Only the top label, bottom label and
	// content region
	// are clickable.
	OnClick *OnClick `json:"onClick,omitempty"`

	// TopLabel: The text of the top label. Formatted text supported.
	TopLabel string `json:"topLabel,omitempty"`

	// ForceSendFields is a list of field names (e.g. "BottomLabel") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "BottomLabel") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *KeyValue) MarshalJSON() ([]byte, error) {
	type NoMethod KeyValue
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type ListMembershipsResponse struct {
	// Memberships: List of memberships in the requested (or first) page.
	Memberships []*Membership `json:"memberships,omitempty"`

	// NextPageToken: Continuation token to retrieve the next page of
	// results. It will be empty
	// for the last page of results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Memberships") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Memberships") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListMembershipsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListMembershipsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type ListSpacesResponse struct {
	// NextPageToken: Continuation token to retrieve the next page of
	// results. It will be empty
	// for the last page of results. Tokens expire in an hour. An error is
	// thrown
	// if an expired token is passed.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Spaces: List of spaces in the requested (or first) page.
	Spaces []*Space `json:"spaces,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "NextPageToken") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "NextPageToken") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListSpacesResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListSpacesResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Membership: Represents a membership relation in Hangouts Chat.
type Membership struct {
	// CreateTime: The creation time of the membership a.k.a the time at
	// which the member
	// joined the space, if applicable.
	CreateTime string `json:"createTime,omitempty"`

	// Member: Member details.
	Member *User `json:"member,omitempty"`

	// Name: Resource name of the membership, in the form
	// "spaces/*/members/*".
	//
	// Example: spaces/AAAAMpdlehY/members/105115627578887013105
	Name string `json:"name,omitempty"`

	// State: State of the membership.
	//
	// Possible values:
	//   "MEMBERSHIP_STATE_UNSPECIFIED" - Default, do not use.
	//   "JOINED" - The user has joined the space.
	//   "INVITED" - The user has been invited, is able to join the space,
	// but currently has
	// not joined.
	//   "NOT_A_MEMBER" - The user is not a member of the space, has not
	// been invited and is not
	// able to join the space.
	State string `json:"state,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "CreateTime") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CreateTime") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Membership) MarshalJSON() ([]byte, error) {
	type NoMethod Membership
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Message: A message in Hangouts Chat.
type Message struct {
	// ActionResponse: Input only. Parameters that a bot can use to
	// configure how its response is
	// posted.
	ActionResponse *ActionResponse `json:"actionResponse,omitempty"`

	// Annotations: Output only. Annotations associated with the text in
	// this message.
	Annotations []*Annotation `json:"annotations,omitempty"`

	// ArgumentText: Plain-text body of the message with all bot mentions
	// stripped out.
	ArgumentText string `json:"argumentText,omitempty"`

	// Cards: Rich, formatted and interactive cards that can be used to
	// display UI
	// elements such as: formatted texts, buttons, clickable images. Cards
	// are
	// normally displayed below the plain-text body of the message.
	Cards []*Card `json:"cards,omitempty"`

	// CreateTime: Output only. The time at which the message was created in
	// Hangouts Chat
	// server.
	CreateTime string `json:"createTime,omitempty"`

	// FallbackText: A plain-text description of the message's cards, used
	// when the actual cards
	// cannot be displayed (e.g. mobile notifications).
	FallbackText string `json:"fallbackText,omitempty"`

	// Name: Resource name, in the form "spaces/*/messages/*".
	//
	// Example: spaces/AAAAMpdlehY/messages/UMxbHmzDlr4.UMxbHmzDlr4
	Name string `json:"name,omitempty"`

	// PreviewText: Text for generating preview chips. This text will not be
	// displayed to the
	// user, but any links to images, web pages, videos, etc. included here
	// will
	// generate preview chips.
	PreviewText string `json:"previewText,omitempty"`

	// Sender: The user who created the message.
	Sender *User `json:"sender,omitempty"`

	// Space: The space the message belongs to.
	Space *Space `json:"space,omitempty"`

	// Text: Plain-text body of the message.
	Text string `json:"text,omitempty"`

	// Thread: The thread the message belongs to.
	Thread *Thread `json:"thread,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "ActionResponse") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ActionResponse") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *Message) MarshalJSON() ([]byte, error) {
	type NoMethod Message
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// OnClick: An onclick action (e.g. open a link).
type OnClick struct {
	// Action: A form action will be trigger by this onclick if specified.
	Action *FormAction `json:"action,omitempty"`

	// OpenLink: This onclick triggers an open link action if specified.
	OpenLink *OpenLink `json:"openLink,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Action") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Action") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *OnClick) MarshalJSON() ([]byte, error) {
	type NoMethod OnClick
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// OpenLink: A link that opens a new window.
type OpenLink struct {
	// Url: The URL to open.
	Url string `json:"url,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Url") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Url") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *OpenLink) MarshalJSON() ([]byte, error) {
	type NoMethod OpenLink
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Section: A section contains a collection of widgets that are
// rendered
// (vertically) in the order that they are specified. Across all
// platforms,
// cards have a narrow fixed width, so
// there is currently no need for layout properties (e.g. float).
type Section struct {
	// Header: The header of the section, text formatted supported.
	Header string `json:"header,omitempty"`

	// Widgets: A section must contain at least 1 widget.
	Widgets []*WidgetMarkup `json:"widgets,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Header") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Header") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Section) MarshalJSON() ([]byte, error) {
	type NoMethod Section
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Space: A room or DM in Hangouts Chat.
type Space struct {
	// DisplayName: Output only. The display name (only if the space is a
	// room).
	DisplayName string `json:"displayName,omitempty"`

	// Name: Resource name of the space, in the form "spaces/*".
	//
	// Example: spaces/AAAAMpdlehYs
	Name string `json:"name,omitempty"`

	// Type: Output only. The type of a space.
	//
	// Possible values:
	//   "TYPE_UNSPECIFIED"
	//   "ROOM"
	//   "DM"
	Type string `json:"type,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "DisplayName") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DisplayName") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Space) MarshalJSON() ([]byte, error) {
	type NoMethod Space
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// TextButton: A button with text and onclick action.
type TextButton struct {
	// OnClick: The onclick action of the button.
	OnClick *OnClick `json:"onClick,omitempty"`

	// Text: The text of the button.
	Text string `json:"text,omitempty"`

	// ForceSendFields is a list of field names (e.g. "OnClick") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "OnClick") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *TextButton) MarshalJSON() ([]byte, error) {
	type NoMethod TextButton
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// TextParagraph: A paragraph of text. Formatted text supported.
type TextParagraph struct {
	Text string `json:"text,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Text") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Text") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *TextParagraph) MarshalJSON() ([]byte, error) {
	type NoMethod TextParagraph
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Thread: A thread in Hangouts Chat.
type Thread struct {
	// Name: Resource name, in the form "spaces/*/threads/*".
	//
	// Example: spaces/AAAAMpdlehY/threads/UMxbHmzDlr4
	Name string `json:"name,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Name") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Name") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Thread) MarshalJSON() ([]byte, error) {
	type NoMethod Thread
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// User: A user in Hangouts Chat.
type User struct {
	// DisplayName: The user's display name.
	DisplayName string `json:"displayName,omitempty"`

	// Name: Resource name, in the format "users/*".
	Name string `json:"name,omitempty"`

	// Type: User type.
	//
	// Possible values:
	//   "TYPE_UNSPECIFIED" - Default value for the enum. DO NOT USE.
	//   "HUMAN" - Human user.
	//   "BOT" - Bot user.
	Type string `json:"type,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DisplayName") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DisplayName") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *User) MarshalJSON() ([]byte, error) {
	type NoMethod User
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// UserMentionMetadata: Annotation metadata for user mentions (@).
type UserMentionMetadata struct {
	// Type: The type of user mention.
	//
	// Possible values:
	//   "TYPE_UNSPECIFIED" - Default value for the enum. DO NOT USE.
	//   "ADD" - Add user to space.
	//   "MENTION" - Mention user in space.
	Type string `json:"type,omitempty"`

	// User: The user mentioned.
	User *User `json:"user,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Type") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Type") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *UserMentionMetadata) MarshalJSON() ([]byte, error) {
	type NoMethod UserMentionMetadata
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// WidgetMarkup: A widget is a UI element that presents texts, images,
// etc.
type WidgetMarkup struct {
	// Buttons: A list of buttons. Buttons is also oneof data and only one
	// of these
	// fields should be set.
	Buttons []*Button `json:"buttons,omitempty"`

	// Image: Display an image in this widget.
	Image *Image `json:"image,omitempty"`

	// KeyValue: Display a key value item in this widget.
	KeyValue *KeyValue `json:"keyValue,omitempty"`

	// TextParagraph: Display a text paragraph in this widget.
	TextParagraph *TextParagraph `json:"textParagraph,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Buttons") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Buttons") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *WidgetMarkup) MarshalJSON() ([]byte, error) {
	type NoMethod WidgetMarkup
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// method id "chat.spaces.get":

type SpacesGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Returns a space.
func (r *SpacesService) Get(name string) *SpacesGetCall {
	c := &SpacesGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SpacesGetCall) Fields(s ...googleapi.Field) *SpacesGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *SpacesGetCall) IfNoneMatch(entityTag string) *SpacesGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *SpacesGetCall) Context(ctx context.Context) *SpacesGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *SpacesGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *SpacesGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "chat.spaces.get" call.
// Exactly one of *Space or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Space.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *SpacesGetCall) Do(opts ...googleapi.CallOption) (*Space, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Space{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns a space.",
	//   "flatPath": "v1/spaces/{spacesId}",
	//   "httpMethod": "GET",
	//   "id": "chat.spaces.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. Resource name of the space, in the form \"spaces/*\".\n\nExample: spaces/AAAAMpdlehY",
	//       "location": "path",
	//       "pattern": "^spaces/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+name}",
	//   "response": {
	//     "$ref": "Space"
	//   }
	// }

}

// method id "chat.spaces.list":

type SpacesListCall struct {
	s            *Service
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists spaces the caller is a member of.
func (r *SpacesService) List() *SpacesListCall {
	c := &SpacesListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	return c
}

// PageSize sets the optional parameter "pageSize": Requested page size.
// The value is capped at 1000.
// Server may return fewer results than requested.
// If unspecified, server will default to 100.
func (c *SpacesListCall) PageSize(pageSize int64) *SpacesListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": A token
// identifying a page of results the server should return.
func (c *SpacesListCall) PageToken(pageToken string) *SpacesListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SpacesListCall) Fields(s ...googleapi.Field) *SpacesListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *SpacesListCall) IfNoneMatch(entityTag string) *SpacesListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *SpacesListCall) Context(ctx context.Context) *SpacesListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *SpacesListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *SpacesListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/spaces")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "chat.spaces.list" call.
// Exactly one of *ListSpacesResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListSpacesResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *SpacesListCall) Do(opts ...googleapi.CallOption) (*ListSpacesResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListSpacesResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists spaces the caller is a member of.",
	//   "flatPath": "v1/spaces",
	//   "httpMethod": "GET",
	//   "id": "chat.spaces.list",
	//   "parameterOrder": [],
	//   "parameters": {
	//     "pageSize": {
	//       "description": "Requested page size. The value is capped at 1000.\nServer may return fewer results than requested.\nIf unspecified, server will default to 100.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A token identifying a page of results the server should return.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/spaces",
	//   "response": {
	//     "$ref": "ListSpacesResponse"
	//   }
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *SpacesListCall) Pages(ctx context.Context, f func(*ListSpacesResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "chat.spaces.members.get":

type SpacesMembersGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Returns a membership.
func (r *SpacesMembersService) Get(name string) *SpacesMembersGetCall {
	c := &SpacesMembersGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SpacesMembersGetCall) Fields(s ...googleapi.Field) *SpacesMembersGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *SpacesMembersGetCall) IfNoneMatch(entityTag string) *SpacesMembersGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *SpacesMembersGetCall) Context(ctx context.Context) *SpacesMembersGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *SpacesMembersGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *SpacesMembersGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "chat.spaces.members.get" call.
// Exactly one of *Membership or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Membership.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *SpacesMembersGetCall) Do(opts ...googleapi.CallOption) (*Membership, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Membership{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns a membership.",
	//   "flatPath": "v1/spaces/{spacesId}/members/{membersId}",
	//   "httpMethod": "GET",
	//   "id": "chat.spaces.members.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. Resource name of the membership to be retrieved, in the form\n\"spaces/*/members/*\".\n\nExample: spaces/AAAAMpdlehY/members/105115627578887013105",
	//       "location": "path",
	//       "pattern": "^spaces/[^/]+/members/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+name}",
	//   "response": {
	//     "$ref": "Membership"
	//   }
	// }

}

// method id "chat.spaces.members.list":

type SpacesMembersListCall struct {
	s            *Service
	parent       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists human memberships in a space.
func (r *SpacesMembersService) List(parent string) *SpacesMembersListCall {
	c := &SpacesMembersListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	return c
}

// PageSize sets the optional parameter "pageSize": Requested page size.
// The value is capped at 1000.
// Server may return fewer results than requested.
// If unspecified, server will default to 100.
func (c *SpacesMembersListCall) PageSize(pageSize int64) *SpacesMembersListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": A token
// identifying a page of results the server should return.
func (c *SpacesMembersListCall) PageToken(pageToken string) *SpacesMembersListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SpacesMembersListCall) Fields(s ...googleapi.Field) *SpacesMembersListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *SpacesMembersListCall) IfNoneMatch(entityTag string) *SpacesMembersListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *SpacesMembersListCall) Context(ctx context.Context) *SpacesMembersListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *SpacesMembersListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *SpacesMembersListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+parent}/members")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "chat.spaces.members.list" call.
// Exactly one of *ListMembershipsResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListMembershipsResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *SpacesMembersListCall) Do(opts ...googleapi.CallOption) (*ListMembershipsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListMembershipsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists human memberships in a space.",
	//   "flatPath": "v1/spaces/{spacesId}/members",
	//   "httpMethod": "GET",
	//   "id": "chat.spaces.members.list",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "pageSize": {
	//       "description": "Requested page size. The value is capped at 1000.\nServer may return fewer results than requested.\nIf unspecified, server will default to 100.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A token identifying a page of results the server should return.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "parent": {
	//       "description": "Required. The resource name of the space for which membership list is to be\nfetched, in the form \"spaces/*\".\n\nExample: spaces/AAAAMpdlehY",
	//       "location": "path",
	//       "pattern": "^spaces/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+parent}/members",
	//   "response": {
	//     "$ref": "ListMembershipsResponse"
	//   }
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *SpacesMembersListCall) Pages(ctx context.Context, f func(*ListMembershipsResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "chat.spaces.messages.create":

type SpacesMessagesCreateCall struct {
	s          *Service
	parent     string
	message    *Message
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Create: Creates a message.
func (r *SpacesMessagesService) Create(parent string, message *Message) *SpacesMessagesCreateCall {
	c := &SpacesMessagesCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.message = message
	return c
}

// ThreadKey sets the optional parameter "threadKey": Opaque thread
// identifier string that can be specified to group messages
// into a single thread. If this is the first message with a given
// thread
// identifier, a new thread is created. Subsequent messages with the
// same
// thread identifier will be posted into the same thread. This relieves
// bots
// and webhooks from having to store the Hangouts Chat thread ID of a
// thread (created earlier by them) to post
// further updates to it.
//
// Has no effect if thread field,
// corresponding to an existing thread, is set in message.
func (c *SpacesMessagesCreateCall) ThreadKey(threadKey string) *SpacesMessagesCreateCall {
	c.urlParams_.Set("threadKey", threadKey)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SpacesMessagesCreateCall) Fields(s ...googleapi.Field) *SpacesMessagesCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *SpacesMessagesCreateCall) Context(ctx context.Context) *SpacesMessagesCreateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *SpacesMessagesCreateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *SpacesMessagesCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.message)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+parent}/messages")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "chat.spaces.messages.create" call.
// Exactly one of *Message or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Message.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *SpacesMessagesCreateCall) Do(opts ...googleapi.CallOption) (*Message, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Message{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a message.",
	//   "flatPath": "v1/spaces/{spacesId}/messages",
	//   "httpMethod": "POST",
	//   "id": "chat.spaces.messages.create",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "parent": {
	//       "description": "Required. Space resource name, in the form \"spaces/*\".\nExample: spaces/AAAAMpdlehY",
	//       "location": "path",
	//       "pattern": "^spaces/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "threadKey": {
	//       "description": "Opaque thread identifier string that can be specified to group messages\ninto a single thread. If this is the first message with a given thread\nidentifier, a new thread is created. Subsequent messages with the same\nthread identifier will be posted into the same thread. This relieves bots\nand webhooks from having to store the Hangouts Chat thread ID of a thread (created earlier by them) to post\nfurther updates to it.\n\nHas no effect if thread field,\ncorresponding to an existing thread, is set in message.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+parent}/messages",
	//   "request": {
	//     "$ref": "Message"
	//   },
	//   "response": {
	//     "$ref": "Message"
	//   }
	// }

}

// method id "chat.spaces.messages.delete":

type SpacesMessagesDeleteCall struct {
	s          *Service
	name       string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Delete: Deletes a message.
func (r *SpacesMessagesService) Delete(name string) *SpacesMessagesDeleteCall {
	c := &SpacesMessagesDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SpacesMessagesDeleteCall) Fields(s ...googleapi.Field) *SpacesMessagesDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *SpacesMessagesDeleteCall) Context(ctx context.Context) *SpacesMessagesDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *SpacesMessagesDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *SpacesMessagesDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "chat.spaces.messages.delete" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *SpacesMessagesDeleteCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Empty{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes a message.",
	//   "flatPath": "v1/spaces/{spacesId}/messages/{messagesId}",
	//   "httpMethod": "DELETE",
	//   "id": "chat.spaces.messages.delete",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. Resource name of the message to be deleted, in the form\n\"spaces/*/messages/*\"\n\nExample: spaces/AAAAMpdlehY/messages/UMxbHmzDlr4.UMxbHmzDlr4",
	//       "location": "path",
	//       "pattern": "^spaces/[^/]+/messages/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+name}",
	//   "response": {
	//     "$ref": "Empty"
	//   }
	// }

}

// method id "chat.spaces.messages.get":

type SpacesMessagesGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Returns a message.
func (r *SpacesMessagesService) Get(name string) *SpacesMessagesGetCall {
	c := &SpacesMessagesGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SpacesMessagesGetCall) Fields(s ...googleapi.Field) *SpacesMessagesGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *SpacesMessagesGetCall) IfNoneMatch(entityTag string) *SpacesMessagesGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *SpacesMessagesGetCall) Context(ctx context.Context) *SpacesMessagesGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *SpacesMessagesGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *SpacesMessagesGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "chat.spaces.messages.get" call.
// Exactly one of *Message or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Message.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *SpacesMessagesGetCall) Do(opts ...googleapi.CallOption) (*Message, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Message{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns a message.",
	//   "flatPath": "v1/spaces/{spacesId}/messages/{messagesId}",
	//   "httpMethod": "GET",
	//   "id": "chat.spaces.messages.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. Resource name of the message to be retrieved, in the form\n\"spaces/*/messages/*\".\n\nExample: spaces/AAAAMpdlehY/messages/UMxbHmzDlr4.UMxbHmzDlr4",
	//       "location": "path",
	//       "pattern": "^spaces/[^/]+/messages/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+name}",
	//   "response": {
	//     "$ref": "Message"
	//   }
	// }

}

// method id "chat.spaces.messages.update":

type SpacesMessagesUpdateCall struct {
	s          *Service
	name       string
	message    *Message
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Update: Updates a message.
func (r *SpacesMessagesService) Update(name string, message *Message) *SpacesMessagesUpdateCall {
	c := &SpacesMessagesUpdateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.message = message
	return c
}

// UpdateMask sets the optional parameter "updateMask": Required. The
// field paths to be updated.
//
// Currently supported field paths: "text", "cards".
func (c *SpacesMessagesUpdateCall) UpdateMask(updateMask string) *SpacesMessagesUpdateCall {
	c.urlParams_.Set("updateMask", updateMask)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SpacesMessagesUpdateCall) Fields(s ...googleapi.Field) *SpacesMessagesUpdateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *SpacesMessagesUpdateCall) Context(ctx context.Context) *SpacesMessagesUpdateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *SpacesMessagesUpdateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *SpacesMessagesUpdateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.message)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "chat.spaces.messages.update" call.
// Exactly one of *Message or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Message.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *SpacesMessagesUpdateCall) Do(opts ...googleapi.CallOption) (*Message, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Message{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates a message.",
	//   "flatPath": "v1/spaces/{spacesId}/messages/{messagesId}",
	//   "httpMethod": "PUT",
	//   "id": "chat.spaces.messages.update",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Resource name, in the form \"spaces/*/messages/*\".\n\nExample: spaces/AAAAMpdlehY/messages/UMxbHmzDlr4.UMxbHmzDlr4",
	//       "location": "path",
	//       "pattern": "^spaces/[^/]+/messages/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "updateMask": {
	//       "description": "Required. The field paths to be updated.\n\nCurrently supported field paths: \"text\", \"cards\".",
	//       "format": "google-fieldmask",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+name}",
	//   "request": {
	//     "$ref": "Message"
	//   },
	//   "response": {
	//     "$ref": "Message"
	//   }
	// }

}
