try
    source $VIM/vimrc 
catch
    echo "\$VIM/vimrc not found"
endtry

set tabstop=4
set shiftwidth=4
set expandtab
syntax on
