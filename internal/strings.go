package main

import (
	"net/http"
)

var banner = `
                                                                               
 _  _  _               __               ____                               __  
| || || |              \ \             |  _ \                             / _) 
| \| |/ |   _  __  ___  \ \     _____  | |_) ) __  ____   ___  _____ _  __\ \  
 \_   _/ | | |/  \/ / |  > \   (_____) |  _ ( /  \/ /\ \ / / |/ / __) |/ / _ \ 
   | | | |_| ( ()  <| | / ^ \          | |_) | ()  <  \ v /|   <> _)| / ( (_) )
   |_| | ._,_|\__/\_\\_)_/ \_\         |____/ \__/\_\  > < |_|\_\___)__/ \___/ 
       | |                                            / ^ \                    
       |_|                                           /_/ \_\                   
`

var messageTemplate = `From: %s
To: %s
Subject: %s

%s

You have received %d/1 emails today. Use the /buy command to purchase a one-time paid plan to support bot development and remove the limit.`

var paidMessageTemplate = `From: %s
To: %s
Subject: %s

%s`

var declinedMessageTemplate = `Message from %s to %s rejected due to free plan limit exceeded.
Use the /buy command to purchase a one-time paid plan to support bot development and remove the limit.`

var internalServerErrorMessage = &HttpError{
	Status:  http.StatusInternalServerError,
	Message: "Internal server error",
}

var notFoundMessage = &HttpError{
	Status:  http.StatusNotFound,
	Message: "The user with the specified ID was not found in the API database. Please register with the bot",
}

var badRequestMessage = &HttpError{
	Status:  http.StatusBadRequest,
	Message: "Bad request. Couldn't parse an ID",
}
