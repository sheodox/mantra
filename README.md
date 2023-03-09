# Mantra

Mantra is an experiment that reminds you occasionally of things you don't want to forget. Use this to give yourself words of encouragement, or remind yourself to do things you want to become habits.

Currently it uses Discord webhooks to send you messages.

## Setup

### Prerequisites

1. Docker
1. Docker Compose

You also need to setup a `.env` file in `src/static/.env`, if you run `./run.sh` first it will create one for you, edit the file according to the comments in the newly created `.env` file.

### Running

1. `./run.sh prod`

### Updating

1. `git pull`
1. `./run.sh prod`
