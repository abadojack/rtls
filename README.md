# RTLS

## How to run
- Configure .env as in sample.env
- Make sure you have `docker` and `docker-compose` installed
- Run `docker-compose up --build`

## Endpoints

#### POST /players
 - Request body
```json
{
    "score": "123"
}
```
- Response
```json
{
    "id": "80bb08a5-0d7d-4907-9576-4bb8cce9b268",
    "score": 123
}
```

#### GET /players/[id]/rank
- Response
```json
{
    "id": "80bb08a5-0d7d-4907-9576-4bb8cce9b268",
    "score": 123,
    "rank": 6
}
```

#### PUT /players/[id]/score
- Request body
```json
{
    "score": "270"
}
```
- Response
```json
{
   "id": "80bb08a5-0d7d-4907-9576-4bb8cce9b268",
   "score": 270
}
```

#### GET /leaderboard/top/[integer]
- Response
```json
[
  {
    "id": "7987aa37-ae91-4927-a5ad-1ccf40694974",
    "score": 300,
    "rank": 1
  },
  {
    "id": "9360adb0-bad1-4e4d-a8fe-9a55f2269336",
    "score": 300,
    "rank": 1
  },
  {
    "id": "80bb08a5-0d7d-4907-9576-4bb8cce9b268",
    "score": 270,
    "rank": 2
  },
  {
    "id": "8e38b101-30ae-4f0e-b715-1746615d841b",
    "score": 250,
    "rank": 3
  },
  {
    "id": "df8d07ca-cc82-4992-bcea-8b3bdc6c952f",
    "score": 250,
    "rank": 3
  }
]
```


## Design choices

- ### packages
- 1. `gorm` - best orm for Golang out there. Saves a lot of dev time and quite efficient
- 2. `gorilla/mux`  - best request router and dispatcher for Golang

- ### Top N players
```
Algorithm GetLeaderBoard(n)
Input: n (integer) - the number of top players to retrieve
    // Check if the requested number of players (n) exceeds the limit in Redis
    if n > limit:
        // Fetch the top n players from the database
        players, err := getNPlayers(n)
        
        // If there's an error fetching players from the database, return the error
        if err is not nil:
            return nil, err
        
        // Attempt to set the fetched players in the Redis leaderboard
        err = setLeaderboardToRedis(players)
        
        // Log the error if setting data in Redis fails, but continue execution
        if err is not nil:
            log_error("error writing leaderboard to redis", err)
        
        // Update the limit and last player based on the fetched players
        limit = length of players
        last = last element of players
        
        // Return the fetched players as the leaderboard
        return players, nil
    end if
    
    // If n does not exceed the limit in Redis, attempt to load players from Redis
    players, err := loadLeaderboardFromRedis()
    
    // If there's an error loading data from Redis, return the error
    if err is not nil:
        return nil, err
    
    // If there are players loaded from Redis
    if length of players > 0:
        // Return the top n players from the loaded data
        return players[:n], nil
    
    // If there are no players loaded from Redis, return nil
    return nil, nil
End Algorithm
```