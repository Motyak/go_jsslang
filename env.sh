[ "${BASH_SOURCE[0]}" == "$0" ] && { >&2 echo "this script is meant to be sourced, not executed"; exit 1; }

function gorun {
    local file="$1"
    local args="${@:2}"

    file="./$file"
    [[ "$file" =~ _main\.go$ ]] || {
        [ -f "${file%.go}_main.go" ] && {
            file="${file%.go}_main.go"
        }
    }

    local dir="${file%/*}/"
    local filename="${file##*/}"

    cmd="GO111MODULE=off go run -tags ${filename%.go} $dir${args:+ }$args"
    ((DRYRUN)) && >&2 echo -n "NOT "
    >&2 echo "Executing cmd \`$cmd\`"
    ((DRYRUN)) && return 0
    eval "$cmd"
}
