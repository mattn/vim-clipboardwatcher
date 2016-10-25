let s:cmd = expand('<sfile>:h:h:gs!\\!/!') . '/clipboardwatcher/clipboardwatcher.exe'
if !filereadable(s:cmd)
  finish
endif

function! clipboardwatcher#watch(cb)
  if exists('s:job')
    call job_stop(s:job)
  endif
  let s:job = job_start(s:cmd, { 'out_cb': a:cb, 'out_mode': 'nl' })
endfunction
