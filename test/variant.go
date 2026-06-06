package main

import "fmt"

type value_t struct {
    variant any
}

type value_t__Bool = bool
type value_t__Float = float64
type value_t__Str = string
type value_t__List = []value_t
type value_t__Map = map[value_t]value_t

type value_t__Type struct {
    id int
}

var (
    Bool = value_t__Type{1}
    Float = value_t__Type{2}
    Str = value_t__Type{3}
    List = value_t__Type{4}
    Map = value_t__Type{5}
)

func (val value_t) is_nil() bool {
    return val.variant == nil
}

func (val value_t) holds(type_ value_t__Type) bool {
    switch val.variant.(type) {
        case value_t__Bool: return type_ == Bool
        case value_t__Float: return type_ == Float
        case value_t__Str: return type_ == Str
        case value_t__List: return type_ == List
        case value_t__Map: return type_ == Map
        default: panic("should not happen")
    }
}

type value_t__Visitor struct {
    Nil func()
    Bool func(value_t__Bool)
    Float func(value_t__Float)
    Str func(value_t__Str)
    List func(value_t__List)
    Map func(value_t__Map)
}

type value_t__VisitorAny struct {
    Nil func() any
    Bool func(value_t__Bool) any
    Float func(value_t__Float) any
    Str func(value_t__Str) any
    List func(value_t__List) any
    Map func(value_t__Map) any
}

func visit (val value_t, visitor value_t__Visitor) {
    if visitor.Nil == nil {panic("Visitor failed to implement Nil()")}
    if visitor.Bool == nil {panic("Visitor failed to implement Bool()")}
    if visitor.Float == nil {panic("Visitor failed to implement Float()")}
    if visitor.Str == nil {panic("Visitor failed to implement Str()")}
    if visitor.List == nil {panic("Visitor failed to implement List()")}
    if visitor.Map == nil {panic("Visitor failed to implement Map()")}
    if val.is_nil() {
        visitor.Nil()
        return
    }
    switch val_ := val.variant.(type) {
        case value_t__Bool: visitor.Bool(val_)
        case value_t__Float: visitor.Float(val_)
        case value_t__Str: visitor.Str(val_)
        case value_t__List: visitor.List(val_)
        case value_t__Map: visitor.Map(val_)
        default: panic("should not happen")
    }
}

func visitAny (val value_t, visitor value_t__VisitorAny) any {
    if visitor.Nil == nil {panic("Visitor failed to implement Nil()")}
    if visitor.Bool == nil {panic("Visitor failed to implement Bool()")}
    if visitor.Float == nil {panic("Visitor failed to implement Float()")}
    if visitor.Str == nil {panic("Visitor failed to implement Str()")}
    if visitor.List == nil {panic("Visitor failed to implement List()")}
    if visitor.Map == nil {panic("Visitor failed to implement Map()")}
    if val.is_nil() {
        return visitor.Nil()
    }
    switch v := val.variant.(type) {
        case value_t__Bool: return visitor.Bool(v)
        case value_t__Float: return visitor.Float(v)
        case value_t__Str: return visitor.Str(v)
        case value_t__List: return visitor.List(v)
        case value_t__Map: return visitor.Map(v)
        default: panic("should not happen")
    }
}

func doit(val value_t) {

    visit(val, value_t__Visitor{
        Nil: func() {
            fmt.Printf("processing value_t__Nil\n")
        },

        Bool: func(bool_ value_t__Bool) {
            fmt.Printf("processing value_t__Bool: %t\n", bool_)
        },

        Float: func(float_ value_t__Float) {
            fmt.Printf("processing value_t__Float: %f\n", float_)
        },

        Str: func(str_ value_t__Str) {
            fmt.Printf("processing value_t__Str: %v\n", str_)
        },

        List: func(list value_t__List) {
            fmt.Printf("processing value_t__List: %s\n", list)
        },

        Map: func(map_ value_t__Map) {
            fmt.Printf("processing value_t__Map: %s\n", map_)
        },
    })
}
