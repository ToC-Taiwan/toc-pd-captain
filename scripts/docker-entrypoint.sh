#!/bin/bash

export $(cat /toc-pd-captain/.env | xargs)

/toc-pd-captain/toc-pd-captain
