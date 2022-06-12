# required

-   ## Go routines
-   ## channels
-   ## select

# Types

-   ## for-select loop
-   ## done channel
-   ## pipelines

## notes

-   buffered Channel: ` channel with limited capacity. Also, when routine sends data to that channel, it can move ahead without waiting for someone to receive that data from channel. It will get blocked only when channel is full, as soon as receiver reads data, it again frees up space and routine continues to run.`
-   unbuffered Channel: ` When unbuffered channel is created, routine sending data to channel will we in hold/wait state until someone has received that data from channel. after data is received, that routine will continue its execution. It Only allows to communicate synchronously`
