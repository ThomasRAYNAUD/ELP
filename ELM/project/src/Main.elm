module Main exposing (..)

-- carrying out the necessary imports
import Array exposing (Array, fromList, get)
import Browser
import Html exposing (Html, div, text, pre, input, button, h1, p, blockquote,ul,li,ol)
import Html.Attributes exposing (placeholder, value, disabled, class, type_, checked)
import Html.Events exposing (onClick, onInput)
import Http exposing (Error)
import Random exposing (int)
import Json.Decode exposing (Decoder, int, string, list, field, map2, map3)


-- MAIN

main : Program () Model Msg
main =
    Browser.element
        { init = init
        , view = view
        , update = update
        , subscriptions = subscriptions
        }


-- MODEL

-- Defines an alias for the application's data model
type alias Model =
    { selected : Maybe String      -- Word selected for the game
    , userInput : String           -- Current user input
    , words : List String          -- List of words loaded from file
    , fileLoadStatus : FileLoadStatus  -- File upload status
    , correctInputEntered : Bool   -- Indicates correct user input
    , result : Result Error (List Quote)  -- Result of the API request, which is browsed with the list that appears in a sub-list ...
    , checkBoxChecked : Bool        -- Indicates whether the checkbox is checked
    , score : Int                  -- User score
    }

-- Defines an alias for browsing the json for the values that we want
type alias Quote =
    { word : String                -- word return with the request
    , meanings : List Meaning
    }

-- Continue in a "duy JSON floor"
type alias Meaning =
    { partOfSpeech : String         -- Verb ? Noun ? 
    , definitions : List Definition -- other list
    }

-- Définit un alias pour représenter la définition d'un mot
type alias Definition =
    { definition : String           -- Definition
    , synonyms : List String        -- synonyms
    , antonyms : List String        -- antonyms
    }

-- state of the file
type FileLoadStatus
    = NotLoaded
    | Loading
    | Loaded String                -- File successfully loaded (contains file contents)
    | LoadError                    -- Error


-- UPDATE
--
-- Defines a message type to represent events and actions in the web app
type Msg
    = FindRandom                         -- Message to trigger a random word search
    | RandomNumber Int                   -- Message containing a randomly generated number
    | InitFileLoad (Result Error String) -- Message resulting from initial file loading
    | FileContentLoaded (Result Error String)  -- Message resulting from content loaded from file
    | FileLoadError                      -- Message indicating an error loading the file
    | UserInput String                   -- Message for user input
    | GenerateNewWord                    -- Message to generate a new word
    | CorrectInputEntered Bool            -- Message indicating whether user input is correct
    | GotQuote (Result Error (List Quote)) -- Message resulting from the reception of the meanings of a word
    | ToggleCheckBox                     -- Message to toggle checkbox status (on/off)


-- Model update function in response to messages (Msg)
update : Msg -> Model -> (Model, Cmd Msg)
update msg model =
    case msg of
        -- When the user requests a random word search
        FindRandom ->
            case List.length model.words of
                0 ->
                    -- If the word list is empty, no action is taken
                    (model, Cmd.none)

                _ -> -- in other cases -> Generates a random index to select a word from the list
                    let
                        randomIndex =
                            Random.generate RandomNumber (Random.int 0 (List.length model.words - 1))
                        -- Updates the model with the selected word and resets certain states
                        newModel =
                            { model | selected = Nothing, userInput = "", correctInputEntered = False }
                    in
                    -- Returns the new model and a command to obtain a random number
                    (newModel, randomIndex)

        -- When a random number is generated
        RandomNumber rn ->
            let
                -- Selects the word corresponding to the random index
                selected =
                    get rn (fromList model.words)
            in
            case selected of
                Just word ->
                    -- If the word is found, updates the model and triggers an API request
                    ( { model | selected = selected, userInput = "", correctInputEntered = False }, fetchApi word )
                Nothing ->
                    -- If word not found -> do nothing
                    (model, Cmd.none)

        -- When the initial loading of the file with words is successful
        InitFileLoad (Ok content) ->
            let
                -- Divides file content into words
                newWords =
                    String.split " " content
                -- Generates a random index to select a word from the new list
                randomIndex =
                    Random.generate RandomNumber (Random.int 0 (List.length newWords - 1))
            in
            -- Updates model with new words and loading status
            ( { model | words = newWords, fileLoadStatus = Loaded content, correctInputEntered = False }, randomIndex )

        -- If the initial file upload fails
        InitFileLoad (Err _) ->
            -- Updates the model with a loading error status
            ( { model | fileLoadStatus = LoadError, correctInputEntered = False }, Cmd.none )

        -- If the file content is loaded successfully
        FileContentLoaded (Ok content) ->
            -- Updates the model by resetting certain states
            ( { model | correctInputEntered = False, checkBoxChecked = False }, Cmd.none )

        -- If the file content fails to load
        FileContentLoaded (Err _) ->
            -- Updates model with API loading error
            ( { model | result = Err (Http.BadBody "Error loading API response."), correctInputEntered = False }, Cmd.none )

        -- If a file loading error is detected
        FileLoadError ->
            -- Updates the model with a file loading error
            ( { model | result = Err (Http.BadBody "Error loading file."), correctInputEntered = False }, Cmd.none )

        -- Cases where the meanings of a word are successfully received
        GotQuote result ->
            -- Update with an API request
            ({ model | result = result }, Cmd.none)

        -- Cas where user input is edit
        UserInput inputWord ->
            let
                -- Checks whether the user input is correct in relation to the selected word
                isCorrect =
                    case model.selected of
                        Just correctWord ->
                            inputWord == correctWord

                        Nothing ->
                            False
            in
            -- Updates the model based on user input
            if isCorrect then
                -- If the user input is corrext, +1 in score
                ( { model | userInput = inputWord, correctInputEntered = isCorrect, score = model.score + 1 }, Cmd.none )
            else
                -- If input is false, update without editing the score
                ( { model | userInput = inputWord, correctInputEntered = isCorrect }, Cmd.none )

        --  Case where new word as to be generate
        GenerateNewWord ->
            if model.correctInputEntered then
                -- If the previous entry was correct, generate a new word and activate the checkbox
                case List.length model.words of
                    0 ->
                        -- word list -> nothing inside
                        (model, Cmd.none)

                    _ ->
                        let
                            -- Generates a random index to select a word from the list
                            randomIndex =
                                Random.generate RandomNumber (Random.int 0 (List.length model.words - 1))
                            -- Updates the model with the new word and resets some states
                            newModel =
                                { model | selected = Nothing, userInput = "", correctInputEntered = False, checkBoxChecked = False }
                        in
                        -- Returns the new model and a command to obtain a random number and activate the checkbox
                        (newModel, Cmd.batch [ randomIndex, Cmd.map (\_ -> ToggleCheckBox) Cmd.none ])
            else
                -- input incorrct -> do nothing
                (model, Cmd.none)

        -- When the status of the correct entry is modified
        CorrectInputEntered isCorrect ->
            -- Met à jour le modèle avec l'état de la saisie correcte
            ( { model | correctInputEntered = isCorrect }, Cmd.none )

        -- When checkbox status is toggled
        ToggleCheckBox ->
            -- Met à jour le modèle en basculant l'état de la case à cocher
            ( { model | checkBoxChecked = not model.checkBoxChecked }, Cmd.none )

-- SUBSCRIPTIONS

subscriptions : Model -> Sub Msg
subscriptions _ =
    Sub.none -- no subscription is defined, so we return Sub.none

-- VIEW

-- Function to define the view (user interface) based on the model
view : Model -> Html Msg
view model =
    -- Use the result of the API request to determine the display
    case model.result of
        -- display the quotes, if OK
        Ok quotes ->
            div []
                [ -- Section for user input and "Next word" button
                  div [ class "user-input-container" ]
                    [ -- Main title (h1) based on the state of the checkbox
                      h1 [] [ text (if model.checkBoxChecked then Maybe.withDefault "No word ......" model.selected else "Find me ...") ]
                    , -- User input field
                      input [ placeholder "Write the word", value model.userInput, onInput UserInput, class "user-input" ] []
                    , -- Button to generate a new word
                      button [ onClick GenerateNewWord
                             -- The button is disabled if the input is incorrect or the word is not selected
                             , disabled (not model.correctInputEntered || model.userInput /= Maybe.withDefault "" model.selected)
                             , class "custom-button"
                             ] [ text "Next word" ]
                    , -- Checkbox to print the word "onClick ToggleCheckBox"
                      input [ type_ "checkbox", checked model.checkBoxChecked, onClick ToggleCheckBox ] []
                    ]
                , -- Section to display the score
                  div [] [ text ("Score: " ++ String.fromInt model.score) ]
                , -- Section to display the quotes
                  div [] (List.map viewQuote quotes)
                ]

        -- If the request fails (Err) > display an error message
        Err _ ->
            div []
                [ text "Error loading quotes" ]

-- Function to render a single quote in the view
viewQuote : Quote -> Html Msg
viewQuote quote =
    div []
        [ 
            -- Render the meanings of the quote
            div [] (List.map viewMeaning quote.meanings)
        ]

-- Function to render a single meaning in the view
viewMeaning : Meaning -> Html Msg
viewMeaning meaning =
    div []
        [ ul []
            [ li [] [ text meaning.partOfSpeech ]
            , -- Render the definitions for the meaning
              ol [] (List.map viewDefinition meaning.definitions)
            ]
        ]

-- Function to render a single definition in the view
viewDefinition : Definition -> Html Msg
viewDefinition definition =
    li []
        [ blockquote []
            [ p [] [ text definition.definition ]
            , -- Render synonyms if they exist
              if List.isEmpty definition.synonyms then
                  text ""
              else
                  p [] [ text "Synonyms: ", text (String.join ", " definition.synonyms) ]
            , -- Render antonyms if they exist
              if List.isEmpty definition.antonyms then
                  text ""
              else
                  p [] [ text "Antonyms: ", text (String.join ", " definition.antonyms) ]
            ]
        ]
-- ul - li - ol > html balise for list (ol number list)


-- INIT

-- Initializer function to set up the initial model and initiate a file load request
init : () -> (Model, Cmd Msg)
init _ =
    ( { initialModel | score = 0 }
    , Http.get
        { url = "static/words.txt"
        , expect = Http.expectString InitFileLoad
        }
    )

-- FETCH API

-- Function to initiate an API request based on a word
fetchApi : String -> Cmd Msg
fetchApi word =
    Http.get
        { url = "https://api.dictionaryapi.dev/api/v2/entries/en/" ++ word
        , expect = Http.expectJson GotQuote (list quoteDecoder)
        }

-- DECODERS

-- JSON decoders for decoding API responses

quoteDecoder : Decoder Quote
quoteDecoder =
    map2 Quote
        (field "word" string)
        (field "meanings" (list meaningDecoder))

meaningDecoder : Decoder Meaning
meaningDecoder =
    map2 Meaning
        (field "partOfSpeech" string)
        (field "definitions" (list definitionDecoder))

definitionDecoder : Decoder Definition
definitionDecoder =
    map3 Definition
        (field "definition" string)
        (field "synonyms" (list string))
        (field "antonyms" (list string))

-- INITIAL MODEL

-- Initial model with default values
initialModel : Model
initialModel =
    { selected = Nothing
    , userInput = ""
    , words = []
    , fileLoadStatus = NotLoaded
    , correctInputEntered = False
    , result = Ok []
    , checkBoxChecked = False
    , score = 0
    }
