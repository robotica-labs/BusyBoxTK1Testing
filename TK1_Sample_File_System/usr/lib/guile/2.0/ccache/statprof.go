GOOF----LE-4-2.0q      ] Ø 4      hq      ] g  guile¤	 ¤	g  define-module*¤	 ¤	 ¤	g  statprof¤	 ¤	g  filenameS¤		f  statprof.scm¤	
g  importsS¤	g  srfi¤	g  srfi-1¤	 ¤	 ¤	g  system¤	g  vm¤	 ¤	 ¤	g  frame¤	 ¤	 ¤	g  program¤	 ¤	 ¤	 ¤	g  exportsS¤	g  statprof-active?¤	g  statprof-start¤	g  statprof-stop¤	g  statprof-reset¤	g  statprof-accumulated-time¤	 g  statprof-sample-count¤	!g  statprof-fold-call-data¤	"g  statprof-proc-call-data¤	#g  statprof-call-data-name¤	$g  statprof-call-data-calls¤	%g  statprof-call-data-cum-samples¤	&g  statprof-call-data-self-samples¤	'g  statprof-call-data->stats¤	(g  statprof-stats-proc-name¤	)g  statprof-stats-%-time-in-proc¤	*g  statprof-stats-cum-secs-in-proc¤	+g   statprof-stats-self-secs-in-proc¤	,g  statprof-stats-calls¤	-g  !statprof-stats-self-secs-per-call¤	.g   statprof-stats-cum-secs-per-call¤	/g  statprof-display¤	0g  statprof-display-anomolies¤	1g  statprof-fetch-stacks¤	2g  statprof-fetch-call-tree¤	3g  with-statprof¤	4g  gcprof¤	5 !"#$%&'()*+,-./01234 ¤	6g  	autoloadsS¤	7g  ice-9¤	8g  format¤	978 ¤	:8 ¤	;9: ¤	<g  set-current-module¤	=< ¤	>< ¤	?g  accumulated-time¤	@g  last-start-time¤	Ag  sample-count¤	Bg  sampling-frequency¤	Cg  remaining-prof-time¤	Dg  profile-level¤	Eg  %count-calls?¤	Fg  gc-time-taken¤	Gg  record-full-stacks?¤	Hg  stacks¤	Ig  procedure-data¤	Jg  make-call-data¤	Kg  call-data-proc¤	Lg  procedure-name¤	Mg  call-data-name¤	Ng  with-output-to-string¤	Og  write¤	Pg  call-data-printable¤	Qg  call-data-call-count¤	Rg  call-data-cum-sample-count¤	Sg  call-data-self-sample-count¤	Tg  inc-call-data-call-count!¤	Ug  inc-call-data-cum-sample-count!¤	Vg   inc-call-data-self-sample-count!¤	Wg  make-syntax-transformer¤	XW ¤	YW ¤	Zg  accumulate-time¤	[g  macro¤	\g  $sc-dispatch¤	]\ ¤	^\ ¤	_g  _¤	`g  any¤	a_`¤	bg  syntax->datum¤	cb ¤	db ¤	eg  datum->syntax¤	fe ¤	ge ¤	hg  set!¤	ig  +¤	je  0.0¤	kg  -¤	l@ ¤	mg  syntax-violation¤	nm ¤	om ¤	pf  -source expression failed to match any pattern¤	qg  program?¤	rg  program-num-free-variables¤	sg  program-objcode¤	tg  	hashq-ref¤	ug  
hashq-set!¤	vg  get-call-data¤	wg  stack-length¤	xg  frame-procedure¤	yg  
count-call¤	zg  frame-previous¤	{g  make-hash-table¤	|g  	hash-fold¤	}g  and=>¤	~g  	stack-ref¤	g  sample-stack-procs¤ g  inside-profiler?¤ g  	positive?¤ g  get-internal-run-time¤ g  
make-stack¤ g  profile-signal-handler¤ g  pk¤ g  what!¤ g  set-vm-trace-level!¤ g  the-vm¤ g  vm-trace-level¤ g  	setitimer¤ g  ITIMER_PROF¤ g  assq¤ g  gc-stats¤ g  	add-hook!¤ g  vm-apply-hook¤ g  remove-hook!¤ g  error¤ f  /Can't reset profiler while profiler is running.¤ g  	sigaction¤ g  SIGPROF¤ f  :Can't call statprof-fold-called while profiler is running.¤ e  100.0¤ e  1.0¤ g  max¤ g  stats-sorter¤ g  current-output-port¤ f  No samples recorded.
¤ g  sort¤ f  !~5a ~10a   ~7a ~8a ~8a ~8a  ~8@a
¤ f  %  ¤ f  
cumulative¤  f  self¤ ¡f   ¤ ¢f  total¤ £f   ~5a  ~9a  ~8a ~8a ~8a ~8a  ~8@a
¤ ¤f  time¤ ¥f  seconds¤ ¦f  calls¤ §f  ms/call¤ ¨f  name¤ ©f  ~5a ~10a   ~7a  ~8@a
¤ ªf  %¤ «f  ~5a  ~10a  ~7a  ~8@a
¤ ¬g  for-each¤ ­f  #~6,2f ~9,2f ~9,2f ~7d ~8,2f ~8,2f  ¤ ®f  ~6,2f ~9,2f ~9,2f  ¤ ¯g  display¤ °g  newline¤ ±f  ---
¤ ²g  simple-format¤ ³f  Sample count: ~A
¤ ´f  *Total time: ~A seconds (~A seconds in GC)
¤ µg  internal-time-units-per-second¤ ¶f  ==[~A ~A ~A]
¤ ·f  Total time: ~A
¤ ¸f  5Can't get accumulated time while profiler is running.¤ ¹g  procedure=?¤ ºg  map¤ »g  lists->trees¤ ¼g  cadr¤ ½g  find¤ ¾g  	assq-set!¤ ¿g  filter¤ Àg  identity¤ Ág  unfold-right¤ Âg  stack->procedures¤ Ãg  loopS¤ ÄÃ¤ Åg  hzS¤ ÆÅ	¤ Çg  count-calls?S¤ ÈÇ	¤ Ég  full-stacks?S¤ ÊÉ	¤ ËÄÆÈÊ ¤ Ìg  inexact->exact¤ Íg  floor¤ Îe  	1000000.0¤ Ïf  Invalid macro body¤ Ðg  keyword?¤ Ñg  eq?¤ Òg  @¤ ÓÒ ¤ Ôg  lambda¤ ÕÉ	¤ ÖÄÕ ¤ ×g  after-gc-hook¤C 5  hX\  F  ]4	
56;5	 4> >  "  G   ?R@RARBRCR
DRER
FRGRHRIR  h   Ë   ]  C  Ã       g  proc
		 g  
call-count		 g  cum-sample-count			 g  self-sample-count			  g  filenamef  statprof.scm
 Ã
	 Ä	 			  g  nameg  make-call-data CJRh   t   ] 
£C l       g  cd
		  g  filenamef  statprof.scm
 Å
	 Å	 		  g  nameg  call-data-proc CKRLK h   }   ]4 56   u       g  cd
		  g  filenamef  statprof.scm
 Æ
	 Æ	,	 Æ	 		  g  nameg  call-data-name CMRMNOK   h   R   ] 4L 56   J       g  filenamef  statprof.scm
 É		 É	/	 É	( 		
   C   h       ]	4 5$  C O 6        g  cd
		 g  t			  g  filenamef  statprof.scm
 Ç
	 È			 È		 É	 		  g  nameg  call-data-printable CPR  h   z   ] £C r       g  cd
		  g  filenamef  statprof.scm
 Ê
	 Ê	" 		  g  nameg  call-data-call-count CQR h      ] 	£Cx       g  cd
		  g  filenamef  statprof.scm
 Ë
	 Ë	( 		  g  nameg  call-data-cum-sample-count CRR   h      ] 	£Cy       g  cd
		  g  filenamef  statprof.scm
 Ì
	 Ì	) 		  g  nameg  call-data-self-sample-count CSR  h      ]  £¤C          g  cd
		  g  filenamef  statprof.scm
 Î
		 Ï		
 Ï		 Ï	 		  g  nameg  inc-call-data-call-count! CTR  h      ] 	 	£¤C        g  cd
		  g  filenamef  statprof.scm
 Ð
	 Ñ		 Ñ		 Ñ	 		  g  nameg  inc-call-data-cum-sample-count! CUR    h      ] 	 	£¤C        g  cd
		  g  filenamef  statprof.scm
 Ò
	 Ó		 Ó		 Ó	 		  g  nameg   inc-call-data-self-sample-count! CVR4YZ[^adgh?ijkl    h    `   ]   C     X       g  	stop-time
		  g  filenamef  statprof.scm
 Õ
	 Ö	 		   C     h   a   ]	4 5L 4?6Y       g  args
		 g  v			  g  filenamef  statprof.scm	 Õ
 		   Cop        h(   _   ]	4 5$   O @ 6 W       g  y
		' g  tmp		'  g  filenamef  statprof.scm
 Õ
 		'   C5ZRqrstIJu  hp   #  ]4 5$  "  	4 5
$   "  4 545$  C4 


54>  "  G  C       g  proc
		o g  t	
	# g  k	4	o g  t		?	o g  	call-data		T	o  g  filenamef  statprof.scm
 Ù
	 Ú		
 Ú		
 Ú		 Û		 Û		' Ú		. Ý		4 Ú		7 Þ		? Þ		K ß		T ß		W à	
 		o  g  nameg  get-call-data CvRwGHAxyz{u|Uv     h      ]4 56          g  proc
		 g  val		 g  accum			  g  filenamef  statprof.scm
 õ			 ÷		 ö	 			   C}vV~    h  ü  ])H4 >  "  G  $    "    "  »$  45$  Z&  K454	5"ÿÿÇ4	>  "  G  45$  "  "ÿÿ45"ÿÿ|4
>  "  G  445>  "  !G  "  4 
54	5"ÿÿ+JC       ô      g  stack
		 g  hit-count-call?		 g  frame		1 ì g  
procs-seen		1 ì g  self		1 ì g  t		> ¸  g  filenamef  statprof.scm
 æ
	 ç		 ç		 ê		" ë		$ ë		+ í		- í		1 ï		7 ò		8 ü		> ò		N þ		Q		R		Y	,	j		k	 	 		  	 ¡	 µ	 ¹ ô	 Î ú	 Ñ ú	 à ú	 ì ï	 í ï	 õ ð	 ï	  		  g  nameg  sample-stack-procs CRRDE?j@B   hð   T  ] 45$  Û45 45$  "  44554	5$  "  7
$  $445 445 5>  "  G  "    4

>  "  G  $  "  145  
$  $445 445 5>  "  G  "   "    CL      g  sig
	 ð g  	stop-time	 ç g  t			= g  stack		= ä g  inside-apply-trap?		F á  g  filenamef  statprof.scm

															.		2		3	#	:		=		@	!	F		N!		X)		Y*		\*	'	a+	+	d+	;	j+	+	k+	'	p*	  ×	 ,	 .	 0	 1	 .	 ©3	 ®5	$ ´5	 º6	 »7	 ¾7	' Ã8	+ Æ8	; Ì8	+ Í8	' Ò7	 î:	 +	 ð  g  nameg  profile-signal-handler CR?j@}xTv  h   e   ]4 56   ]       g  proc
		  g  filenamef  statprof.scm
E		G		F	 		   C 	       h@   »   ]$  C45  44 5>  "  G  45  C    ³       g  frame
		<  g  filenamef  statprof.scm
?
	@		 ×		B		B		D		D		+D		4I		:I	 		<  g  nameg  
count-call CyRD  h   ð   ] 6è       g  filenamef  statprof.scm
M
	P	 		
  g  nameg  statprof-active?g  documentationf  uReturns @code{#t} if @code{statprof-start} has been called more times
than @code{statprof-stop}, @code{#f} otherwise. CRDC@FFBEy   hè   g  ] $  Ñ  $  %4 5$  "  4 5"   45  445 5 	$  4


  >  "  G  "  4


>  "  G  $   4445 5>  "  G  "   445 445 5>  "  G  CC      _      g  rpt
	 à g  t	!	: g  use-rpt?	? à  g  filenamef  statprof.scm
S
	W		W		X		X		Y		Z		[	 	[	+	![	 	![		0\	 	5\	+	7\	 	?Y		D]		E^		K^		L`		P`		Q`	(	W`		X`		Z_		`a		ab		jb	'	mb	1	rb		c	 e	 f	 c	 g	 h	 ¢h	 ¥h	& «h	 ²h	 ¿i	 Âi	 Çi	* Êi	: Ði	* Ñi	& Öi	 /	 â
  g  nameg  statprof-startg  documentationf  Start the profiler.@code{} CRDFFEyC?j@    h     ]  
$  z445 5 445 445 5>  "  G  	$   4
445 5>  "  G  "   4



5 45   CC           g  filenamef  statprof.scm
m
	q		q		
r		r		u		u		u	+	u		u		u		!t		"v		%v		*v	*	-v	:	3v	*	4v	&	9v		Gw		Hx		Kx		Nx	)	Tx		[x		h{	"	t{		x ×		y|	 |	 }	  	 
  g  nameg  statprof-stopg  documentationf  Stop the profiler.@code{} CRDE?@ABC{IGH      hx   Û  - . , 3 #  45$  4>  "  G  "    
  
   	 
4 5   4>  "  G  C    Ó      g  sample-seconds
		t g  sample-microseconds		t g  count-calls?			t g  full-stacks?			t  g  filenamef  statprof.scm

			 		!		%		*		:		=		@		C		H		J		M		N		W		[		\		^		_	 		t	  g  nameg  statprof-resetg  documentationf VReset the statprof sampler interval to @var{sample-seconds} and
@var{sample-microseconds}. If @var{count-calls?} is true, arrange to
instrument procedure calls as well as collecting statistical profiling
data. If @var{full-stacks?} is true, collect all sampled stacks into a
list for later analysis.

Enables traps and debugging as necessary. CRD|     h      ]L 6             g  key
		
 g  value		
 g  prior-result			
  g  filenamef  statprof.scm
¡		
¢	 		
	   CI   h8     ]45$  4>  "  G  "    O 6            g  proc
		2 g  init		2  g  filenamef  statprof.scm

											2 	 		2	  g  nameg  statprof-fold-call-datag  documentationf 7Fold @var{proc} over the call-data accumulated by statprof. Cannot be
called while statprof is active. @var{proc} should take two arguments,
@code{(@var{call-data} @var{prior-result})}.

Note that a given proc-name may appear multiple times, but if it does,
it represents different functions with the same name. C!RDtI  h0     ]45$  4>  "  G  "    6           g  proc
		+  g  filenamef  statprof.scm
¦
	©		©		ª		ª		ª		+¬	 		+  g  nameg  statprof-proc-call-datag  documentationf  TReturns the call-data associated with @var{proc}, or @code{#f} if
none is available. C"RPSR E$j     h¨   ¢  ]14 54 54 545 45 45 $  4 5"  		$  
$  
"  	"  $  	45"   C         g  	call-data
	 ¥ g  	proc-name		 ¥ g  self-samples		 ¥ g  cum-samples		 ¥ g  all-samples		" ¥ g  secs-per-sample		/ ¥ g  	num-calls		C ¥  g  filenamef  statprof.scm
±
	»			»		¼		»		½		»		¾		"»		%¿		*À		/¿		/»		7Á		8Á	'	C»		LÄ		NÄ	,	OÄ		UÅ	+	VÅ		]Æ	,	^Æ		gÈ		jÉ		oÉ		qÉ	*	zÊ		|Ê	9	Ê	 Ë	 Ì	 Í	 Ð	 Ì	 ¤Ã	 %	 ¥  g  nameg  statprof-call-data->statsg  documentationf  0Returns an object of type @code{statprof-stats}. C'R h      ] 
£C y       g  stats
		  g  filenamef  statprof.scm
Ò
	Ò	) 		  g  nameg  statprof-stats-proc-name C(R  h      ] £C ~       g  stats
		  g  filenamef  statprof.scm
Ó
	Ó	. 		  g  nameg  statprof-stats-%-time-in-proc C)R     h      ] 	£C       g  stats
		  g  filenamef  statprof.scm
Ô
	Ô	0 		  g  nameg  statprof-stats-cum-secs-in-proc C*R   h      ] 	£C       g  stats
		  g  filenamef  statprof.scm
Õ
	Õ	1 		  g  nameg   statprof-stats-self-secs-in-proc C+R  h   }   ] 	£Cu       g  stats
		  g  filenamef  statprof.scm
Ö
	Ö	% 		  g  nameg  statprof-stats-calls C,R      h      ] 	£C       g  stats
		  g  filenamef  statprof.scm
×
	×	2 		  g  nameg  !statprof-stats-self-secs-per-call C-R h      ] 	£C       g  stats
		  g  filenamef  statprof.scm
Ø
	Ø	1 		  g  nameg   statprof-stats-cum-secs-per-call C.R+*  h8   ã   ]
4 545
$  4 545"  6    Û       g  x
		4 g  y		4 g  diff			4  g  filenamef  statprof.scm
Ü
	Ý		
Þ		Ý		Ý		à			à		á		%â		,á			4ß	 		4	  g  nameg  stats-sorter CR 8!' h   ~   ]4 5C   v       g  data
		 g  prior-value		  g  filenamef  statprof.scm
ï		ð	 	ð	 			   CE ¡¢£¤¥¦§¨©ª«¬E8­)*+,-.®¯(°      h    =  ]$  J4M 4 54 54 54 5è4 5è4	 5>  "  G  "  )4M 
4 54 54 5>  "  G  44 5M >  "  G  M 6      5      g  stats
	   g  filenamef  statprof.scm
õ		ö			÷		÷		ø		ù		ú		%û		/ü		6ü		:ý		Aý		F÷		Sþ		Yþ		Zÿ		a 		h		sþ		|			 	 	 	   g  nameg  display-stats-line C¯±²³´Fµ $     h(    -  1  3  H J (  45 K "   45 
$  J 64545	$  H4J 
>	  "  G  4J >	  "  G  "  84J >  "  G  4J >  "  G  4 O >  "  G  4J >  "  G  445 >  "  G  4 5 !"#6         g  port
	% g  
stats-list	9% g  sorted-stats		D%  g  filenamef  statprof.scm
å
	è		è		è		ë		$ë		)ê		/ì		1ì		2î		7ò		9î		9î		<ó		Dî		L		M		S		U		W		Y	(	[	/	]	2	_	9	a	A	f		o			u			w
		y
		{
	&	}
	0	
	8 
	B 
	L 		 	 	 	 	 ¡	& £	- ¨	 ±	 ·	 ¹	 »	 ½	& ¿	0 Ä	 Í	 æ	 ê	 ñ	 ú	 ÿ	 	-					&"	%	 @		%


  g  nameg  statprof-displayg  documentationf  Displays a gprof-like summary of the statistics collected. Unless an
optional @var{port} argument is passed, uses the current output port. C/R!EQR²¶M      hH   á   ]$  =4 5
$  .44 55$  4 54 54 56CCC Ù       g  data
		G g  prior-value		G  g  filenamef  statprof.scm
												 		 		! 		%			*"		+#		2$		9%		A!		 		G	   C²·³   h@     ] 4>  "  G  445 >  "  G  45 6             g  filenamef  statprof.scm

			'		'		'	'	%'		2(		3(	)	9(	 
		9
  g  nameg  statprof-display-anomoliesg  documentationf  QA sanity check that attempts to detect anomolies in statprof's
statistics.@code{} C0RD¸?µ        h0   ò   ] 45$  4>  "  G  "   C       ê       g  filenamef  statprof.scm
*
	,		,		-		-		-		(.	 		)
  g  nameg  statprof-accumulated-timeg  documentationf  AReturns the time accumulated during the last statprof run.@code{} CRD¸A     h(   ì   ] 45$  4>  "  G  "   C  ä       g  filenamef  statprof.scm
0
	2		2		3		3		3	 		&
  g  nameg  statprof-sample-countg  documentationf  HReturns the number of samples taken during the last statprof run.@code{} C RMi#RQi$RRi%RSi&RH       h   >  ] C   6      g  filenamef  statprof.scm
;
 		
  g  nameg  statprof-fetch-stacksg  documentationf  ÇReturns a list of stacks, as they were captured since the last call
to @code{statprof-reset}.

Note that stacks are only collected if the @var{full-stacks?} argument
to @code{statprof-reset} is true. C1Rqs       h8   È   ] &  C4 5$  45$  4 545CCC  À       g  a
		6 g  b		6  g  filenamef  statprof.scm
D		
E		F		G		E		G		"G		#H		*H		1H	 		6	  g  nameg  procedure=? C¹Rº»    h      ] 4 L 5C       x       g  tail
		  g  filenamef  statprof.scm
R		S	 	T	 	T	.	T	 	S	 		   Ci¼h   z   ] C    r       g  a
		 g  b		  g  filenamef  statprof.scm
X		X	%		X	.	X	" 			   C½  h   k   ]L L 6   c       g  x
		  g  filenamef  statprof.scm
[		[	 	
[	(	[	 		   C¾ 
       h¸     ]""   (  +4O 5445?45C(  "ÿÿ¶4O 5$  !4	5"ÿÿ} "ÿÿ` 
"ÿÿR          g  lists
	 ´ g  equal?	 ´ g  in		 ¦ g  
n-terminal		 ¦ g  tails		 ¦ g  trees			7 g  t		` ¦  g  filenamef  statprof.scm
N
	O		P		R		R		V		"V	"	,V		-W		6V		:Y		>P		AZ	
	DZ		PZ		Q[		`P		k^		n`		ua		xb	 	|b	*	}b		`	 ^	 d	
 f	 f	! f	 f	
 ¦d	 ¦O	 ªO	, ´O	 #	 ´	  g  nameg  lists->trees C»R¿ÀÁ     h   Y   ] C  Q       g  x
		  g  filenamef  statprof.scm
j		j	$ 		   Cxz~      h       ]44 
556            g  stack
		  g  filenamef  statprof.scm
h
	j	
	m		j	
	i	 		  g  nameg  stack->procedures CÂR»ºÂH¹  h   T  ] 4455C   L      g  filenamef  statprof.scm
o
	v		v		v		v	 		
  g  nameg  statprof-fetch-call-treeg  documentationf  ¶Return a call tree for the previous statprof run.

The return value is a list of nodes, each of which is of the type:
@code
 node ::= (@var{proc} @var{count} . @var{nodes})
@end code C2RË    h8      ]"  # 
$  C4L >   "  G    "ÿÿÝL "ÿÿÕ              g  i
		)  g  filenamef  statprof.scm
											#		)		)	 			1
   CÌÍÎ   h@   Ù   ] 444L554L4L55LL >  "  G  6   Ñ       g  filenamef  statprof.scm
								&		-		&						)		0		0	#	7	%	0	&	-	'	&	)		2		>	 		>
   C/I     h(   [   ] 4>   "  G  4>   "  G   C S       g  filenamef  statprof.scm
						%	 		'
   CÌÍÎ       h@   Ù   ] 444L554L4L55LL >  "  G  6   Ñ       g  filenamef  statprof.scm
								&		-		&						)		0		0	#	7	%	0	&	-	'	&	)		2		>	 		>
   C/I     h(   [   ] 4>   "  G  4>   "  G   C S       g  filenamef  statprof.scm
						%	 		'
   C/I   h(   [   ] 4>   "  G  4>   "  G   C S       g  filenamef  statprof.scm
						%	 		'
   C  h      - /   0   3 #  #  	d#  #  O  Q O 4O >   "  G  V4>   X4>   "  G  CX4>   "  G  F     	      g  thunk
	  g  loop	  g  hz		  g  count-calls?		  g  full-stacks?		  g  thunk		A   g  filenamef  statprof.scm
x
	A	 	 
g  loopSg  hzS	g  count-calls?S	g  full-stacks?S	   g  nameg  statprofg  documentationf ØProfiles the execution of @var{thunk}.

The stack will be sampled @var{hz} times per second, and the thunk itself will
be called @var{loop} times.

If @var{count-calls?} is true, all procedure calls will be recorded. This
operation is somewhat expensive.

If @var{full-stacks?} is true, at each sample, statprof will store away the
whole call tree, for later analysis. Use @code{statprof-fetch-stacks} or
@code{statprof-fetch-call-tree} to retrieve the last-stored stacks. CR4Y3[^adgÏÐÑ      hP   	  ](  645$    &  C  "ÿÿÈ4 5$  CC        g  kw
		N g  args		N g  def			N  g  filenamef  statprof.scm
°		±		²		²		³		³		³		±		´		#´		&µ	
	-¶		:¶	
	;·		H±	 		N	  g  nameg  
kw-arg-ref CÓÔÃÅÇÉ        h`   Î   -  1  3 O Q 4 54 54 	d54 54 5 
C     Æ       g  args
			[ g  
kw-arg-ref		[  g  filenamef  statprof.scm

	º		»		%º		)¼		4º		5½	
	Aº		B¾		Mº		N¿		Zº	 			[


   C       h   a   ]	4 5L 4?6Y       g  args
		 g  v			  g  filenamef  statprof.scm	
 		   Cop        h(   S  ]	4 5$   O @ 6 K      g  y
		' g  tmp		'  g  filenamef  statprof.scm

 		'  g  documentationf §Profiles the expressions in its body.

Keyword arguments:

@table @code
@item #:loop
Execute the body @var{loop} number of times, or @code{#f} for no looping

default: @code{#f}
@item #:hz
Sampling rate

default: @code{20}
@item #:count-calls?
Whether to instrument each function call (expensive)

default: @code{#f}
@item #:full-stacks?
Whether to collect away all sampled stacks into a list

default: @code{#f}
@end tableg  
macro-typeg  defmacrog  defmacro-argsg  args  C53RÖ?j@ 
      hx   .  ]  $   C 45 4L 
5  $   "  4455  4>  "  G   	 45  	  C    &      g  t
		t g  t
	!	? g  	stop-time
	?	o g  stack	?	o  g  filenamef  statprof.scm
Ú		Û		Þ		á		å		!å		0æ		4æ		5æ	"	<æ		?á		Dç		Y ×		bè		cé		ié		rë	 		t
  g  nameg  gc-callback CD?@AE{IGHCFF×        h      ] 45$  4>  "  G  "   
  
  4	 5 
L   $  O 45  445 5 4L >  "  G  445 445 5>  "  G  CC  y      g  filenamef  statprof.scm
		Ð		Ð		Ñ		Ñ		Ñ		&Ò		)Ó		,Ô		/Õ		0Ö		9Ö		=×		>Ø		@Ø		Cî		Eî		Iï		Mï		Pñ	
	Qò	 	Wò	
	Xó	#	\ó	)	]ó	8	có	#	dó		fó	
	gô	
	{õ	
	~õ	 õ	, õ	< õ	, õ	( õ	
 %	 
   C     h8      ]"  # 
$  C4L >   "  G    "ÿÿÝL "ÿÿÕ              g  i
		)  g  filenamef  statprof.scm
									
		#		)		)	 			1
   CDFF×?j@/I   hh   â   ]  
$  =445 5 4L >  "  G  	4
5   "   4>   "  G   C       Ú       g  filenamef  statprof.scm
		ù		ù		
ú		ú		ý		ý		ý	-	ý		ý		ý		!ü	
	"þ	
	9 ×		:ÿ		Eÿ	
	H 	
	M		_	 		a
   C   h    W  - /   0   3 #  #  O Q O O O Q  Q Q 4>   "  G  V4>   X4>   "  G  CX4>   "  G  FO      g  thunk
	   g  loop	   g  full-stacks?		   g  gc-callback		-   g  pre		L   g  thunk		L   g  post		L    g  filenamef  statprof.scm
Á
	;	 	  
g  loopSg  full-stacks?S	   g  nameg  gcprofg  documentationf +Do an allocation profile of the execution of @var{thunk}.

The stack will be sampled soon after every garbage collection, yielding
an approximate idea of what is causing allocation in your program.

Since GC does not occur very frequently, you may need to use the
@var{loop} parameter, to cause @var{thunk} to be called @var{loop}
times.

If @var{full-stacks?} is true, at each sample, statprof will store away the
whole call tree, for later analysis. Use @code{statprof-fetch-stacks} or
@code{statprof-fetch-call-tree} to retrieve the last-stored stacks. C4RC    >      g  m
		0  g  filenamef  statprof.scm		}
	4 ¯
	8 °
	< ±
	@ ²
	D ³
	H ´
	L µ
	P ¶
	T ¸
	U º		X º
	\ ¿
F Ã
Ï Å
p Æ
´ Ç
E Ê
Û Ë
t Ì
$ Î
Ú Ð
 Ò
 Ù
 æ


.?
;M
ÂS
$m
®
!¶
#¦
&±
'$Ò
'ÁÓ
(cÔ
)Õ
)Ö
*=×
*ÜØ
,Ü
3Åå
6
7Å*
8÷0
8þ6
97
98
99
:q;
;C
@BN
Ah
Co
KÜx
\RÁ
 ;	\T
   C6 