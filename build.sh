#!/bin/sh
echo "🎨 Compiling assets..."
sass sass/main.scss:static/css/main.css
templ generate

