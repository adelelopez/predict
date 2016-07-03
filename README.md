# predict
Calibrate your predictions. Tighten your debugging feedback loop. See how underconfident or overconfident you are. 

## CLI:

### Basic usage: 


predict "\<name>" \<probability> 
-- creates a new prediction with \<name> and \<probability> 

predict judge \<outcome>
-- judges the most recent unjudged prediction with outcome \<outcome>

predict history
-- outputs the prediction history, with most recent ones first

predict stats
-- displays the user's Brier score and a graph 

predict last
-- (TODO) displays the most recent prediction made

predict help
-- (TODO) displays a help menu/blurb

### Advanced usage:

predict set \<option>=\<value>
-- (TODO) set options

predict get "<name>"
-- (TODO) displays information about the specified prediction, disambiguates if necessary 

predict edit "<name>"
-- (TODO) allows the user to edit the specified prediction, disambiguating if necessary
-- to judge you must use 'predict judge'

predict delete "<name>"
-- (TODO) allows the user to delete the specified prediction, disambiguating if necessary

predict judge "<name>" <outcome>
-- (TODO) allows the user to judge the specified prediction, disambiguating if necessary

### (TODO) Options: 

-t \<tag>: tags predictions with \<tag>, or only considers predictions with tag \<tag>
-u 		: restricts history to unjudged predictions
-j 		: restricts history to judged predictions
-v 		: displays version

default-tag : which tag to use by default
verbose 	: how much stuff gets printed
mirror 		: bucket p and 1-p together

## Examples: (These are mostly my vision for how I want it to work, and do not yet work like this.)

```
>> predict “the problem is with opening the file” 60% -t debugging
3/26/15 11:32pm: ”the problem is with opening the file” 	60% -		tagged: debugging

>> predict set default-tag=debugging

>> predict judge true
3/26/15 11:32pm: ”the problem is with opening the file” 	60% true	tagged: debugging

>> predict "the build will succeed" .9 
3/26/15 11:46pm: ”the build will succeed” 					90% -		tagged: debugging

>> predict "docker logs are overfull" 87 
3/26/15 11:46pm: ”docker logs are overfull” 				87% -		tagged: debugging

>> predict hist
3/26/15 11:46pm: ”docker logs are overfull” 				87% -		tagged: debugging
3/26/15 11:46pm: ”the build will succeed” 					90% -		tagged: debugging
3/26/15 11:32pm: ”the problem is with opening the file” 	60% true	tagged: debugging
3/26/15 11:28pm: ”the build will succeed” 					70% false	tagged: debugging

>> predict hist -u
3/26/15 11:46pm: ”docker logs are overfull” 				87% -		tagged: debugging
3/26/15 11:46pm: ”the build will succeed” 					90% -		tagged: debugging
	
>> predict last
3/26/15 11:46pm: ”docker logs are overfull” 				87% -		tagged: debugging

>> predict judge ”docker logs are overfull” true
3/26/15 11:46pm: ”docker logs are overfull” 				87% true	tagged: debugging

>> predict judge "the build will succeed" true
1. 3/26/15 11:46pm: ”the build will succeed” 				90% -		tagged: debugging
2. 3/26/15 11:28pm: ”the build will succeed” 				70% false	tagged: debugging
Select prediction to judge: 1
3/26/15 11:46pm: ”the build will succeed” 					90% true	tagged: debugging

>> predict stats -t debugging
Score: 0.038
::::=1=3=5==10===15===20===25=======33==============50===============66======75===80===85===90===95===99
 1% [  * ]                      
 3%  [ *  ]    
 5%    [    * ] 
10%       [     * ]
15%       *     [       ]
20%                [    ]     *
25%                        [*]
33%                               [     ]*
50%                                           [  *            ]
66%                                                               [    * ]
75%                                                 *                      [     ]
80%                                                                             [   ]               *
85%                                                                                   [*  ]
90%                                                                                 [          *     ]
95%                                                                                            [*  ]    
97%                                                                                               [ * ]
99%                                                                                                 [ *]
::::=1=3=5==10===15===20===25=======33==============50===============66======75===80===85===90===95===99
```
