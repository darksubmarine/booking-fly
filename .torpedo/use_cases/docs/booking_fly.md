
Given a frequent flyer user should be able to do a booking fly from our well known fly routes, selecting the
departure airport and the arrival airport, also setting up the from-to fly dates. If the booking is successful, so the
system should calculate the user awards and upgrade it following the rules below:

      - IF the user.Plan is GOLD 
          AND the user accumulated miles (current user miles + trip.Miles) are greater than 1500
          AND the trip.Miles are greater than 2000
        THEN
          user.Miles = current user miles + trip.Miles + 1000
        ELSE
          user.Miles = current user miles + trip.Miles + 200
      
      - IF the user.Plan is SILVER 
          AND the user accumulated miles (current user miles + trip.Miles) are greater than 8000
        THEN
          user.Miles = current user miles + trip.Miles + 500
          user.Plan = GOLD
      
      - IF the user.Plan is BRONZE 
          AND the user accumulated miles (current user miles + trip.Miles) are greater than 4000
        THEN
          user.Miles = current user miles + trip.Miles + 300
          user.Plan = SILVER