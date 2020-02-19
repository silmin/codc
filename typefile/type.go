package typefile

import (
    "time"
)

type Figure struct {
    ID              string           `json:"_id"`
    Areas           []Areas          `json:"areas"`
    CounterHistory  []CounterHistory `json:"counterHistory"`
    CounterSummary  []CounterSummary `json:"counterSummary"`
    DateEnd         time.Time        `json:"dateEnd"`
    DateStart       time.Time        `json:"dateStart"`
    Filename        string           `json:"filename"`
    TrackerSummary  TrackerSummary   `json:"trackerSummary"`
    VideoResolution VideoResolution  `json:"videoResolution"`
}
type XBounds struct {
    XMax float64 `json:"xMax"`
    XMin float64 `json:"xMin"`
}
type Computed struct {
    A       float64 `json:"a"`
    B       float64 `json:"b"`
    XBounds XBounds `json:"xBounds"`
}
type Point1 struct {
    X int `json:"x"`
    Y int `json:"y"`
}
type Point2 struct {
    X int `json:"x"`
    Y int `json:"y"`
}
type RefResolution struct {
    H int `json:"h"`
    W int `json:"w"`
}
type Location struct {
    Point1        Point1        `json:"point1"`
    Point2        Point2        `json:"point2"`
    RefResolution RefResolution `json:"refResolution"`
}
type Areas struct {
    Color    string   `json:"color"`
    Computed Computed `json:"computed"`
    ID       string   `json:"id"`
    Location Location `json:"location"`
    Name     string   `json:"name"`
}
type CounterHistory struct {
    AreaID    string    `json:"area_id"`
    AreaName  string    `json:"area_name"`
    FrameID   int       `json:"frameId"`
    ID        int       `json:"id"`
    Name      string    `json:"name"`
    Timestamp time.Time `json:"timestamp"`
}
type CounterSummary struct {
    Total    int    `json:"_total"`
    AreaID   string `json:"areaId"`
    AreaName string `json:"areaName"`
    Car      int    `json:"car"`
    Bus      int    `json:"bus,omitempty"`
    Truck    int    `json:"truck,omitempty"`
    Person   int    `json:"person,omitempty"`
}
type TrackerSummary struct {
    TotalItemsTracked int `json:"totalItemsTracked"`
}
type VideoResolution struct {
    H int `json:"h"`
    W int `json:"w"`
}
