module Test exposing (..)

import Browser
import Html exposing (..)
import Html.Attributes exposing (value)
import Html.Events exposing (onInput)
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
  | Success (List Word)
    { enteredWord : String }

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
  = GotWord (Result Http.Error (List Word))
  | CheckWord String

update : Msg -> Model -> (Model, Cmd Msg)
update msg model =
    case msg of
        GotWord (Ok words) ->
            (Success { words = words, enteredWord = model.enteredWord }, Cmd.none)
        GotWord (Err error) ->
            (Failure, Cmd.none)
        CheckWord entered ->
            (Success { words = model.words, enteredWord = entered }, Cmd.none)

-- SUBSCRIPTIONS
subscriptions : Model -> Sub Msg
subscriptions model =
  Sub.none

-- VIEW
view : Model -> Html Msg
view model =
  div []
    [ h2 [] [ text "Word Definitions" ]
    , input [ type_ "text", onInput CheckWord] []
    , viewWord model
    ]

viewWord : Model -> Html Msg
viewWord model =
  case model of
    Failure ->
      div []
        [ text "I could not load the word for some reason. "        ]
    Loading ->
      text "Loading..."
    Success { words, enteredWord }->
        div [] (List.map viewWordMeaning words)
    Success { enteredWord } ->
        case words of
            word::_ ->
                case (word.word == enteredWord) of
                    True -> text "You found the word!"
                    False -> text "Not found"
            _ -> text ""

viewWordMeaning : Word -> Html Msg
viewWordMeaning word =
    div []
        [ 
           ul [] (List.map viewMeaning word.meanings)
        ]

viewMeaning : Meaning -> Html Msg
viewMeaning meaning =
    li []
        [ text meaning.partOfSpeech        , ul [] (List.map viewDefinition meaning.definitions)
        ]

viewDefinition : Definition -> Html Msg
viewDefinition def =
    li [] [ text def.definition ]

-- HTTP

getWord : Cmd Msg
getWord =
    Http.get
    { url = "https://api.dictionaryapi.dev/api/v2/entries/en/hello"
    , expect = Http.expectJson GotWord descriptionDecoder
    }
--Decoder
descriptionDecoder : Decoder (List Word)
descriptionDecoder = Json.Decode.list wordDecoder
wordDecoder : Decoder Word
wordDecoder =
    map2 Word
        (field "word" string)
        (field "meanings" (Json.Decode.list meaningDecoder))
meaningDecoder : Decoder Meaning
meaningDecoder =
    map2 Meaning
        (field "partOfSpeech" string)
        (field "definitions" (Json.Decode.list definitionDecoder))
definitionDecoder : Decoder Definition
definitionDecoder =
    Json.Decode.map Definition
        (field "definition" string)
