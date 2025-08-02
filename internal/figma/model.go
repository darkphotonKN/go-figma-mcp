package figma

import "time"

// File represents a Figma file
type File struct {
	Key          string    `json:"key"`
	Name         string    `json:"name"`
	Thumbnail    string    `json:"thumbnail_url,omitempty"`
	LastModified time.Time `json:"last_modified"`
	Version      string    `json:"version"`
}

// FileResponse represents the response from the Figma API when getting a file
type FileResponse struct {
	Document     Document `json:"document"`
	Components   map[string]Component `json:"components"`
	SchemaVersion int     `json:"schemaVersion"`
	Styles       map[string]Style `json:"styles"`
}

// Document represents the root document of a Figma file
type Document struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Children []Node `json:"children"`
}

// Node represents a node in the Figma document tree
type Node struct {
	ID               string                 `json:"id"`
	Name             string                 `json:"name"`
	Type             string                 `json:"type"`
	Visible          bool                   `json:"visible"`
	Children         []Node                 `json:"children,omitempty"`
	BackgroundColor  []float64             `json:"backgroundColor,omitempty"`
	Fills            []Paint               `json:"fills,omitempty"`
	Strokes          []Paint               `json:"strokes,omitempty"`
	StrokeWeight     float64               `json:"strokeWeight,omitempty"`
	CornerRadius     float64               `json:"cornerRadius,omitempty"`
	AbsoluteBoundingBox *Rectangle         `json:"absoluteBoundingBox,omitempty"`
	Constraints      *LayoutConstraint     `json:"constraints,omitempty"`
	Effects          []Effect              `json:"effects,omitempty"`
	Characters       string                `json:"characters,omitempty"`
	Style            *TypeStyle            `json:"style,omitempty"`
}

// Component represents a reusable component
type Component struct {
	Key         string `json:"key"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Style represents a style definition
type Style struct {
	Key         string    `json:"key"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	StyleType   string    `json:"styleType"`
}

// Paint represents fill or stroke paint
type Paint struct {
	Type     string  `json:"type"`
	Color    *Color  `json:"color,omitempty"`
	Opacity  float64 `json:"opacity,omitempty"`
	ImageRef string  `json:"imageRef,omitempty"`
}

// Color represents an RGBA color
type Color struct {
	R float64 `json:"r"`
	G float64 `json:"g"`
	B float64 `json:"b"`
	A float64 `json:"a"`
}

// Rectangle represents a bounding rectangle
type Rectangle struct {
	X      float64 `json:"x"`
	Y      float64 `json:"y"`
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
}

// LayoutConstraint represents layout constraints
type LayoutConstraint struct {
	Vertical   string `json:"vertical"`
	Horizontal string `json:"horizontal"`
}

// Effect represents visual effects like shadows
type Effect struct {
	Type    string  `json:"type"`
	Visible bool    `json:"visible"`
	Radius  float64 `json:"radius,omitempty"`
	Color   *Color  `json:"color,omitempty"`
	Offset  *Vector `json:"offset,omitempty"`
}

// Vector represents a 2D vector
type Vector struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

// TypeStyle represents text styling
type TypeStyle struct {
	FontFamily string  `json:"fontFamily"`
	FontSize   float64 `json:"fontSize"`
	FontWeight int     `json:"fontWeight"`
	LineHeight string  `json:"lineHeightPx"`
	LetterSpacing float64 `json:"letterSpacing"`
}

// Team represents a Figma team
type Team struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Project represents a Figma project
type Project struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// GetFileRequest represents a request to get a file
type GetFileRequest struct {
	FileKey string `json:"file_key" binding:"required"`
	Version string `json:"version,omitempty"`
	IDs     string `json:"ids,omitempty"`
	Depth   int    `json:"depth,omitempty"`
}

// GetImageRequest represents a request to get images
type GetImageRequest struct {
	FileKey string `json:"file_key" binding:"required"`
	IDs     string `json:"ids" binding:"required"`
	Scale   string `json:"scale,omitempty"`
	Format  string `json:"format,omitempty"`
	UseAbsoluteBounds bool `json:"use_absolute_bounds,omitempty"`
}

// ImageResponse represents the response from the images endpoint
type ImageResponse struct {
	Err    *string           `json:"err"`
	Images map[string]string `json:"images"`
}

// CommentRequest represents a request to get comments
type CommentRequest struct {
	FileKey string `json:"file_key" binding:"required"`
}

// Comment represents a comment on a Figma file
type Comment struct {
	ID        string    `json:"id"`
	Message   string    `json:"message"`
	FileKey   string    `json:"file_key"`
	ParentID  string    `json:"parent_id"`
	User      User      `json:"user"`
	CreatedAt time.Time `json:"created_at"`
	ResolvedAt *time.Time `json:"resolved_at"`
	ClientMeta ClientMeta `json:"client_meta"`
}

// User represents a Figma user
type User struct {
	ID     string `json:"id"`
	Handle string `json:"handle"`
	ImgURL string `json:"img_url"`
	Email  string `json:"email"`
}

// ClientMeta represents client metadata for comments
type ClientMeta struct {
	X *float64 `json:"x"`
	Y *float64 `json:"y"`
	NodeID []string `json:"node_id"`
}

// CommentsResponse represents the response from the comments endpoint
type CommentsResponse struct {
	Comments []Comment `json:"comments"`
}