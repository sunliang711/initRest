#!/bin/bash
rpath="$(readlink ${BASH_SOURCE})"
if [ -z "$rpath" ];then
    rpath=${BASH_SOURCE}
fi
root="$(cd $(dirname $rpath) && pwd)"
cd "$root"
shellHeaderLink='https://pic711.oss-cn-shanghai.aliyuncs.com/sh/shell-header.sh'
if [ -e /etc/shell-header.sh ];then
    source /etc/shell-header.sh
else
    (cd /tmp && wget -q "$shellHeaderLink") && source /tmp/shell-header.sh
fi
# write your code below
usage(){
    cat<<-EOF
	Usage: $(basename $0) [projectName]

	projectName has two pattern:
	    1. myproject
	    2. <git repo>/<user>/myproject
	EOF
exit 1
}

if [ -n $GOPATH ];then
    destPrefix=$GOPATH/src
else
    destPrefix=$HOME/go/src
fi

install(){
    projectName=${1}
    if [ -z "$projectName" ];then
        echo -n "Enter your new project name: "
        read projectName
        if [ -z "$projectName" ];then
            echo "Error: empty project name, quit."
            exit 1
        fi
    fi
    if echo "${projectName}" | grep -qE '^[^/]+/[^/]+/[^/]+$' || echo "${projectName}" | grep -qE '^[^/]+$';then
        dest="${destPrefix}/${projectName}"
    else
        cat<<-EOF
		projectName has two pattern:
		    1. myproject
		    2. <git repo>/<user>/myproject
		EOF
        exit 1
    fi
    echo "Destination project dir is: \"${green}${dest}${green}\""
    if [ -e "$dest" ];then
        echo "${dest} already exists."
        exit 1
    fi
    echo "${green}Creating ${dest} ...${reset}"
    mkdir -p "${dest}"
    cd "${dest}"
    git init .>/dev/null

    cat<<-EOF>.gitignore
		.DS_Store
		*.swp
		.idea/
	EOF

    rsync -a "${root}/rootfs/" ${dest}

    mv cmd/PROJECT_NAME "cmd/${projectName}"

    ( find . -iname "*.go" -print0 | xargs -0 perl -i -pe "s|<PROJECT_NAME>|${projectName}|" )
    ( go mod init ${projectName} && go mod tidy )

    # at last initial git commit
    ( git add . && git commit -m 'initial commit: this commit is commited by initgorestserver script' )

}

case $1 in
    -h|--help)
        usage
        ;;
    *)
        install "$@"
        ;;
esac
