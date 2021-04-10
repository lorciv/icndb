# ICNDb API

Short description of the Internet Chuck Norris Database (ICNDb) API as reported on the official web page http://www.icndb.com/api/.

## Joke

The basic unit of information of the ICNDb API is the joke. A joke is made of 3 fields:
- id (int)
- joke (string), containing the actual joke
- categories (array of strings)

For example, here is one:

```json
{
    "id": 531,
    "joke": "Jesus can walk on water, but Chuck...",
    "categories": ["nerdy"]
}
```

By default, the joke's text is HTML encoded. In order to display special characters correctly, the client should take care of decoding the text upon retrieval.

## API Response

The API responds to client requests with a JSON structure. Regardless of the specific request, such JSON structure will always have the following fields:
- type (string)
- value (any JSON value)

In case of sucess, type contains "success" and value contains a JSON value with the actual response. The type of the JSON value depends on the request.

```json
{
    "type": "success",
    "value": (a JSON value)
}
```

For example, when fetching a joke, the value field will contain the actual joke.

In case of error, type contains something other than "success" and value contains an error message describing what happened. For example, when trying to fetch a joke with incorrect ID (see methods below):
```json
{
    "type": "NoSuchQuoteException",
    "value": "No quote with id=12345."
}
```

## Methods

The API is available at `https://api.icndb.com`. Use HTTP GET requests to retrieve information.

### Random joke

Path: `/jokes/random`

Return a random joke.

Response value: a Joke object.

### N random jokes

Path: `/jokes/random/N`
- `N` is a number

Return N random jokes.

Response value: an array of N Joke objects.

### Fetch joke

Path: `/jokes/ID`
- `ID` is a number

Return a specific joke.

Response value: a Joke object.

### Count jokes

Path: `/jokes/count`

Return the number of available jokes.

Response value: an integer.

### Categories

Path: `/categories`

Return all the available categories.

Response value: an array of strings.

### All jokes

Path: `/jokes`

Return all the jokes.

Response value: a (long) array of Joke objects.

## Parameters

### Limiting categories

When fetching random jokes, the paramater `limitTo` can be used to restrict the scope to one or more categories.

```
?limitTo=cat1
?limitTo=[cat1,cat2,...]
```

Similarly, the parameter `exclude` allows to exclude jokes belonging to the given set of categories.

### Changing the name of the main character

It is possible to change the first and last names of the main character when fetching jokes.

```
?firstName=Santa
?lastName=Claus
```
