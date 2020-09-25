A leaderboard server. It will have two API calls:

1) set a new integer score for a player
2) request the n players currently at rank k:k+n (you are free to use a naïve algorithm)

may be required for sqlite/grpc:

export CGO_ENABLED=1
export GO111MODULE=on