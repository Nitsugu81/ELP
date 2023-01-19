module Test exposing (..)

import Browser
import Html exposing (..)
import Html.Attributes exposing (style)
import Html.Events exposing (..)
import Http
import Json.Decode exposing (Decoder, map2, list, field, string)

-- MAIN

main =
  Browser.element
    { init = init
    , update = update
    , subscriptions = subscriptions
    , view = view
    }

-- MODEL

type Model
  = Failure
  | Loading
  | Success Word

type alias Description = List Word

type alias Word =
    { word : String
    , meanings : List Meaning
    }

type alias Meaning =
    { partOfSpeech : String
    , definitions : List Definition
    }

type alias Definition =
    { definition : String
    }

init : () -> (Model, Cmd Msg)
init _ =
  (Loading, getWord)

-- UPDATE

type Msg
  = GotWord (Result Http.Error Word)

update : Msg -> Model -> (Model, Cmd Msg)
update msg model =
  case msg of
    GotWord result ->
      case result of
        Ok word ->
          (Success word, Cmd.none)

        Err _ ->
          (Failure, Cmd.none)

-- SUBSCRIPTIONS

subscriptions : Model -> Sub Msg
subscriptions model =
  Sub.none

-- VIEW

view : Model -> Html Msg
view model =
  div []
    [ h2 [] [ text "Word Definitions" ]
    , viewWord model
    ]

viewWord : Model -> Html Msg
viewWord model =
  case model of
    Failure ->
      div []
        [ text "I could not load the word for some reason. "
        ]

    Loading ->
      text "Loading..."

    Success word ->
      div []
        [ h3 [] [ text word.word ]
        , ul [] (List.map viewDef word.meanings)
        ]

viewDef : Meaning -> Html Msg
viewDef def =
  li []
    [ text def.partOfSpeech
    , ul [] (List.map viewDefinition def.definitions)
    ]

viewDefinition : Definition -> Html Msg
viewDefinition def =
  li [] [ text def.definition ]

-- HTTP

getWord : Cmd Msg
getWord =
  Http.get
    { url = "https://api.dictionaryapi.dev/api/v2/entries/en/hello"
    , expect = Http.expectJson GotWord wordDecoder
    }
--Decoder
descriptionDecoder : Decoder (List Word)
descriptionDecoder = Json.Decode.list wordDecoder 

wordDecoder : Decoder Word
wordDecoder = 
    map2 Word 
        (field "word" string)
        (field "meanings" <| Json.Decode.list meaningDecoder)

meaningDecoder : Decoder Meaning
meaningDecoder = 
    map2 Meaning
        (field "partOfSpeech" string)
        (field "definitions" <| Json.Decode.list definitionDecoder)

definitionDecoder : Decoder Definition
definitionDecoder = 
    Json.Decode.map Definition 
        (field "definition" string)



