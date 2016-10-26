let s:cmd = expand('<sfile>:h:h:gs!\\!/!') . '/clipboardwatcher/clipboardwatcher.exe'
if !filereadable(s:cmd)
  finish
endif

let s:template = {}

function! s:template.stop()
  call job_stop(self.job)
endfunction

function! clipboardwatcher#watch(cb)
  let ctx = copy(s:template)
  let ctx['dir'] = a:dir
  let ctx['job'] = job_start(s:cmd, { 'out_cb': a:cb, 'out_mode': 'json' })
  return ctx
endfunction
