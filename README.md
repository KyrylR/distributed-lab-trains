# Distributed Lab - Trains

An assignment from Distributed Lab

## Description

Given a timetable for crossing trains between several stations in the following format
`train number; departure station; arrival station; cost; departure time; arrival time; arrival time`
( guaranteed no crossings longer than one day). 
It is necessary to get the "best" options (several, if possible) to travel between all stations so 
that each station can be visited 1 time. The requests for the best options:
* Best price
* Best in time

## Explanation of my approach

The solution to this problem is divided into several steps:
1. Parsing the csv file.
   1. Regex is used to clear the file of all characters that are not digits or the following characters: `;`, `:`, `.`
   2. An array of Train structures is created, where the columns in the file correspond to the fields in that structure.
   3. Parsing itself.
   4. If it fails, the program informs you and terminates its work.
2. Then we find unique stations, which are present in the given information and provide this information in the 
following form: `(map[int][]Train), key - statinId`
3. Use goroutines to accelerate the BuildPathsGo function; at this point, a pathTree is generated. 
The pathTree looks like a graph. Structure itself looks like this:
[Link to PathTree struct]()
After creating the structure, we get the same PathTree as an array by using additional functions and start 
building all possible paths using the BuildNewWaysFromPathTree function.
4. The algorithm for building all paths is as follows, using an array of PathTree structures, new possible 
paths are created and, if they are unique to the paths array, they are added to the array. 
In the following iterations, each path is looped through, and if the last station in the path matches the corresponding 
PathTree station, a new path is created, and a copy is created to be able to handle alternate paths.
5. Once we have found all the possible paths, we start analysing them to make them look like they can be conveniently 
recorded. Additional structures have been created for this purpose:
[Link to QueryWay struct]()
6. Querying by cost is relatively easy: we sort all the paths by price, find the cheapest, and return them.
7. Making a query by time is more complicated, so the following structures have been created to make it easier:

[Link to waitToTrain struct]()

[Link to SortTrains struct]()

8. We find the best route and from station to station, then sort them by time and try to match them with the same 
structures, if this completes we get the best route, if not, we take the fastest trains, combine them and return them.
9. At the end we get all the paths ready to be entered. For query by time, we sort the array using GetLowestTime
function and return the result, similarly for query by price, but we sort using GetLowestCost function.

## Executing program
*You should be in the project root folder in order to run the following commands*

The program is invoked according to the scheme:
```
go run ./trains
```

## Run tests

The program is invoked according to the scheme:
```
go test -v ./tests
```

## Authors

ex. Kyrylo Riabov [Gmail](kyryl.ryabov@gmail.com)

## License

This project is licensed under the [MIT] License - see the LICENSE.md file for details
