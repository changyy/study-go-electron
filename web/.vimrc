try
    source $VIM/vimrc 
catch
    echo "\$VIM/vimrc not found"
endtry

set tabstop=2
set shiftwidth=2
set expandtab
syntax on
