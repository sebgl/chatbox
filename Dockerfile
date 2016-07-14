FROM alpine:3.3

COPY chatbox /usr/local/bin/chatbox

ENTRYPOINT ["/usr/local/bin/chatbox"]