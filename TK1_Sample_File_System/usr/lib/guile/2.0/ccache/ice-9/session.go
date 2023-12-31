GOOF----LE-4-2.0xa      ] Ç 4       h;      ] g  guile¤	 ¤	g  define-module*¤	 ¤	 ¤	g  ice-9¤	g  session¤	 ¤		g  filenameS¤	
f  ice-9/session.scm¤	g  importsS¤	g  documentation¤	 ¤	 ¤	g  regex¤	 ¤	 ¤	g  rdelim¤	 ¤	 ¤	g  match¤	 ¤	 ¤	 ¤	g  exportsS¤	g  help¤	g  add-value-help-handler!¤	g  remove-value-help-handler!¤	g  add-name-help-handler!¤	g  remove-name-help-handler!¤	g  apropos-hook¤	 g  apropos¤	!g  apropos-internal¤	"g  apropos-fold¤	#g  apropos-fold-accessible¤	$g  apropos-fold-exported¤	%g  apropos-fold-all¤	&g  source¤	'g  arity¤	(g  procedure-arguments¤	)g  module-commentary¤	* !"#$%&'() ¤	+g  set-current-module¤	,+ ¤	-+ ¤	.g  object-documentation¤	/g  *value-help-handlers*¤	0g  delete!¤	1g  or-map¤	2g  try-value-help¤	3g  *name-help-handlers*¤	4g  try-name-help¤	5g  make-syntax-transformer¤	65 ¤	75 ¤	8g  macro¤	9g  $sc-dispatch¤	:9 ¤	;9 ¤	<g  _¤	=g  any¤	><=¤	?g  syntax->datum¤	@? ¤	A? ¤	Bg  datum->syntax¤	CB ¤	DB ¤	Eg  length¤	Fg  	provided?¤	Gg  display¤	Hg  help-doc¤	Ig  simple-format¤	Jf  ^~A$¤	Kg  regexp-quote¤	Lg  symbol->string¤	Mg  string?¤	Ng  and-map¤	Og  symbol?¤	Pg  quote¤	Qg  
write-line¤	Rf   commentary:¤	Sf  No ~A found for ~A
¤	Tg  
commentary¤	Ug  
help-usage¤	Vg  unquote¤	Wg  
module-ref¤	Xg  current-module¤	Yg  search-documentation-files¤	Zg  begin¤	[Z ¤	\f  ]`help' depends on the `regex' feature.
You don't seem to have regular expressions installed.
¤	]g  syntax-violation¤	^] ¤	_] ¤	`f  -source expression failed to match any pattern¤	ag  map¤	bg  reverse¤	cg  string-append¤	df  /¤	eg  %search-load-path¤	fg  in-vicinity¤	gg  module-filename¤	hg  file-commentary¤	ig  
procedure?¤	jf  a procedure¤	kf  	an object¤	lf  Documentation from file:¤	mf  Did not find any object ¤	nf  named `~A'
¤	of  matching regexp "~A"
¤	pg  for-each¤	qf  ~S: ~S
¤	rg  module-name¤	sg  cadr¤	tg  caddr¤	uf  !`~S' is ~A in the ~S module.

~A
¤	vg  cadddr¤	wf  Documentation found for:
¤	xg  newline¤	yf  No documentation found for:
¤	zf vUsage: (help NAME) gives documentation about objects named NAME (a symbol)
       (help REGEXP) ditto for objects with names matching REGEXP (a string)
       (help 'NAME) gives documentation for NAME, even if it is not an object
       (help ,EXPR) gives documentation for object returned by EXPR
       (help (my module)) gives module commentary for `(my module)'
       (help) gives this text

`help' searches among bindings exported from loaded modules, while
`apropos' searches among bindings visible from the "current" module.

Examples: (help help)
          (help cons)
          (help "output-string")

Other useful sources of helpful information:

(apropos STRING)
(arity PROCEDURE)
(name PROCEDURE-OR-MACRO)
(source PROCEDURE-OR-MACRO)

Tools:

(backtrace)				;show backtrace from last error
(debug)					;enter the debugger
(trace [PROCEDURE])			;trace procedure (no arg => show)
(untrace [PROCEDURE])			;untrace (no arg => untrace all)

(OPTIONSET-options 'full)		;display option information
(OPTIONSET-enable 'OPTION)
(OPTIONSET-disable 'OPTION)
(OPTIONSET-set! OPTION VALUE)

where OPTIONSET is one of debug, read, eval, print

¤	{g  	make-hook¤	|g  run-hook¤	}g  string-length¤	~f  Empty string not allowed¤	g  make-regexp¤ g  module-uses¤ g  
duplicates¤ g  member¤ g  shadow¤ g  value¤ g  full¤ g  module-obarray¤ g  hash-for-each¤ g  regexp-exec¤ f  : ¤ f  	(unbound)¤ f  	 shadowed¤ g  make-hash-table¤ g  	hash-fold¤ g  hashq-get-handle¤ g  
hashq-set!¤ g  hash-create-handle!¤ g  make-fold-modules¤ g  identity¤ g  
submodules¤ g  resolve-module¤ g  root-modules¤ g  hash-map->list¤ g  module-submodules¤ g  module-public-interface¤ g  procedure-source¤ g  macro?¤ g  macro-transformer¤ g  procedure-property¤ g  arglist¤ f   ¤ f  required¤  f   arguments: ¤ ¡f   argument: ¤ ¢f  ', `¤ £f  ' and `¤ ¤f  ', the rest in `¤ ¥f  , ¤ ¦f  optional¤ §f  keyword¤ ¨f  other keywords allowed¤ ©f  the rest in `¤ ªf  '¤ «g  procedure-minimum-arity¤ ¬f   or more¤ ­f   required and ¤ ®f  	 optional¤ ¯f  	 argument¤ °f  
 arguments¤ ±f  .
¤ ²g  error¤ ³² ¤ ´² ¤ µf  no matching pattern¤ ¶g  required¤ ·g  number?¤ ¸g  	make-list¤ ¹g  optional¤ ºg  keyword¤ »g  allow-other-keys?¤ ¼g  rest¤ ½g  system¤ ¾g  vm¤ ¿g  program¤ À½¾¿ ¤ Ág  program?¤ ÂÀÁ ¤ ÃÀÁ ¤ Äg  program-arguments-alist¤ ÅÀÄ ¤ ÆÀÄ ¤C 5hðK  ^  ]4	
*5 4- >  "  G   .      h   r   ]6j       g  name
		 g  value		  g  filenamef  ice-9/session.scm
	$			%	 			   C /R/  h   Ñ  ]  C     É      g  proc
		  g  filenamef  ice-9/session.scm
	'
		.				.	 		  g  nameg  add-value-help-handler!g  documentationf 1Adds a handler for performing `help' on a value.

`proc' will be called as (PROC NAME VALUE). `proc' should return #t to
indicate that it has performed help, a string to override the default
object documentation, or #f to try the other handlers, potentially
falling back on the normal behavior for `help'. CR0/    h   Ö   ]4 5 C Î       g  proc
		  g  filenamef  ice-9/session.scm
	0
		2			2	 		  g  nameg  remove-value-help-handler!g  documentationf  3Removes a handler for performing `help' on a value. CR1  h   _   ] LL 6      W       g  proc
		
  g  filenamef  ice-9/session.scm
	5	
	
	5	 		
   C/    h      ] O 6              g  name
		 g  value		  g  filenamef  ice-9/session.scm
	4
		5	 			  g  nameg  try-value-help C2R3R3       h   F  ]  C     >      g  proc
		  g  filenamef  ice-9/session.scm
	:
		D				D	 		  g  nameg  add-name-help-handler!g  documentationf §Adds a handler for performing `help' on a name.

`proc' will be called with the unevaluated name as its argument. That is
to say, when the user calls `(help FOO)', the name is FOO, exactly as
the user types it.

`proc' should return #t to indicate that it has performed help, a string
to override the default object documentation, or #f to try the other
handlers, potentially falling back on the normal behavior for `help'. CR03       h   Ô   ]4 5 C Ì       g  proc
		  g  filenamef  ice-9/session.scm
	F
		H			H	 		  g  nameg  remove-name-help-handler!g  documentationf  2Removes a handler for performing `help' on a name. CR1    h   _   ] L 6W       g  proc
		  g  filenamef  ice-9/session.scm
	K	
		K	 		   C3    h   x   ] O 6 p       g  name
		  g  filenamef  ice-9/session.scm
	J
		K	 		  g  nameg  try-name-help C4R478;>ADEF4GHIJKLMNOP)QRSTUV2WXY[\        h   "  -  1  3 4 5$ r45$ R 45$  !&  "  4>  "  G  " $  *444	4
555>  "  G  " æ45$  4>  "  G  " Ã$ ¬"  ²45$  (  "  $  g45$  :4>  "  G  4>  "  G  4>  " ?G  " 84>  " $G  " 4>   " G  " 4>   "  ùG  "  ò45	$  à&  g4445 55$  !&  "  4>  "  G  "   4>  "  G  "  l&  _$  N45$  4>  "  G  "   4>  "  G  "  "ÿþl"  "ÿþd"  "ÿþ\"  4>   "  G  C4>  "  G  C4>   "  G  C           g  exp
		 g  name	$s g  t		-p g  t	 ß? g  doc	á g  x	ÂÞ g  t	 A g  x	">  g  filenamef  ice-9/session.scm
	P
	
	S			S			S			V			V			V		!	S		$	[		$	[			'	b		-	_		<	c		A	c	1	Y	f		]	_		^	g		c	h		h	i		i	j		l	j	&	t	j		v	h		{	g	 	m	 	_	 	n	 ­	q	 ±	_	 ¶ 	 Â 	 Ð 	 Ò 	' Ó 	 Ô 	 Ø 	 Ù 	 ß 	 è 	 ú 	( þ 	4 	( 	"	]	'	]	.) 	$0	]	@ 	g	_	h	r	q	r	u	q	x	s	z	s	"~	q		t		t	(	u	(	u	4	v	4	u	(	t		t		w	©	y	Á	w	:Â	w	 Å	]	Ê	]	.Ì	w	+Ó	]	è	~	ê	~	"î	|	ñ		ó		÷	|	ø 	ý 	0  	  	! 	3" 	%	]	*	]	., 	$3	]	^ 	r 	t	W		x	W	}	W			Y			T			U		 `		


   C   h   e   ]	4 5L 4?6]       g  args
		 g  v			  g  filenamef  ice-9/session.scm		P
 		   C_`    h(   ç   ]	4 5$   O @ 6 ß       g  y
		' g  tmp		'  g  filenamef  ice-9/session.scm
	P
 		'  g  documentationf  7(help [NAME])
Prints useful information.  Try `(help)'.g  
macro-typeg  defmacrog  defmacro-argsg  args  C5RaLbccd   h   i   ] 6      a       g  elt
		
  g  filenamef  ice-9/session.scm
 		 	4	
 	! 		
   Cef        hH   x  ])4 54545445?456      p      g  name
		B g  name		B g  reverse-name			B g  leaf			B g  dir-hint-module-name		#	B g  dir-hint		3	B  g  filenamef  ice-9/session.scm
 
	 		 		 		 		 		 		 		! 	(	# 		# 		& 		) 		3 		3 		8 		B 	 		B  g  nameg  module-filename CgRgh     h       ]	4 5$  6C              g  name
		 g  t			  g  filenamef  ice-9/session.scm
 
	  				  	 		  g  nameg  module-commentary C)R"2ijk h0   Ñ   ] 4545$  "   C       É       g  module
		) g  name		) g  object			) g  data			)  g  filenamef  ice-9/session.scm
 ¤		 §	-	 ¨	4	 ¨	-	 ©	4	" «	4	% ¥	'	( ¥	! 			)	   C$YQlGmInopIqrstuv hh   T  ]	44 54 554 5$  3MN44 54 54 54 55MNCM N C      L      g  entry
		b g  entry-summary		b  g  filenamef  ice-9/session.scm
 º		 »	.	 ¼	2		 ½	/	 ½	<	 ½	/	 ¾	/	 »	.	 »		 ¿		& ¿		+ Â	'	- Á	!	. Å	-	3 Æ	1	4 Ç	.	; È	.	B É	.	G É	;	I É	.	J Ê	.	R Å	-	U Å	'	W Ä	!	^ Í	%	` Ì	 		b   CEwG  h   b   ] 6Z       g  entry
		  g  filenamef  ice-9/session.scm
 Ö		 Ö	- 		   CxG      h(   }   ]M $  N "  4>   "  G   6   u       g  entry
		%  g  filenamef  ice-9/session.scm
 Ú		 Û		 Ü		 Ý		% Þ	 		%   CxyG h   b   ] 6Z       g  entry
		  g  filenamef  ice-9/session.scm
 ç		 ç	- 		   C   h@  Ù  ]245(  O4 5$  4>  "  G  64>  "  G  	 $  
"   6HHHH4O >  "  G  J(  "  4J5$  "  J$  -4>  "  G  4J>  "  G  K"   4O J>  "  G  J(  CJ$  K"  4>   "  G  4>  "  G  J6 Ñ      g  term
	? g  regexp	? g  entries		? g  t			e g  first?		i? g  undocumented-entries		i? g  documented-entries		i? g  documentations		i? g  t	 ¢ ¶  	g  filenamef  ice-9/session.scm
 £
	 ¤		 ­		 ¤		 ¤		 ´		 é		 ´		& ë		* ë		/ ë		= ì		> ï		B ï		G ï		U ñ		Y ñ		[ ò		a ó		e ð		g ¶	'	h ·	%	i ¸	!	i µ		r º	  Ñ	  Ò	 ¢ Ò	 ¢ Ò	 ² Ó	 ³ Ó	 º Ñ	 » Õ	 ¿ Õ	 Ä Õ	 Í Ö	 ã Ø	 è Ú	 á	 ã	 ä	 å	& æ	* æ	/ æ	? ç	 /	?	  g  nameg  help-doc CHRGz    h   m   ] 6e       g  filenamef  ice-9/session.scm
 ö
	 ÷		 ÷	 		
  g  nameg  
help-usage CUR4{i	5R|X}~rprLGiWXx        h  Ã  ]4L4 55$  ï4L >  "  G  4>  "  G  4 >  "  G  $  Q45$  "  M$  (4	>  "  G  4>  "  G  "   "  $4	>  "  G  4>  "  G  M$  34L 5445  5&  "  4	>  "  G  "   
6 C   »      g  symbol
	 g  variable	 g  val		T  g  t		]	o  g  filenamef  ice-9/session.scm
>		?		?	*	?		?		@		'A		+A	 	0A		9B		MC		QC		TD	)	TD		WE	+	]E	'	sE	 	tF	' G	' £I	 µJ	 ¹J	' ¾J	 ÌK	 ÍL	* ÖM	* ÙM	6 áM	* åK	 êN	 îN	$ óN	O	 "		   C  h0   ­   ]4 54 5L  LLO 6       ¥       g  module
		) g  name			) g  obarray			)  g  filenamef  ice-9/session.scm
9			:			:		;		:		)=	 		)   C        h°   ¬  - 1 3 445  >  "  G  4 5
$  C4 5445 545 (  "  45	$  "  4
5H4
5H4
5$  
KK"   O 6¤      g  rgx
		 ° g  options		 ° g  match		9 ° g  uses		E ° g  modules		p ° g  shadow		{ ° g  value	  °  g  filenamef  ice-9/session.scm
%
	
'		'		'		#(		*(		/(		1)		3*		9*		<+		?+	 	E+		E*		H,		R-	 	X.	*	].	7	_.	*	a/	*	b.	%	f-		i0	 	p,		p*		s3		w3		{3		{*		~4	 4	 4	 *	 5	 5	 5	 5	 6	 7	 °8	 )		 °
  g  nameg  aproposg  documentationf  CSearch for bindings: apropos regexp {options= 'full 'shadow 'value} C R"    h      ]C       g  module
		 g  name		 g  var			 g  data			  g  filenamef  ice-9/session.scm
U		V	 			   C#X h   á   ] 445 56   Ù       g  rgx
		  g  filenamef  ice-9/session.scm
S
	W		
Y		Y	)	Y		U	 		  g  nameg  apropos-internalg  documentationf  +Return a list of accessible variable names. C!R|XL  hX     ]$  K4L 4 55$  14L 5$  C4L >  "  G  LL 6CC 
      g  name
		W g  var		W g  data			W g  val			T  g  filenamef  ice-9/session.scm
z		{			{		|	0	|		s	 	s	3	s	 	 s		!t	%	-s		1v		Qw	 		W	  g  nameg  module-filter C h(      ] $  L LL O 4 56C          g  module
		% g  data		%  g  filenamef  ice-9/session.scm
p		~			(	"~	 		%	  g  nameg  fold-module C      h@   î  ]445 >  "  G  4545  O 6    æ      g  proc
		< g  init		< g  rgx			< g  folder			< g  match		'	< g  recorded		'	<  g  filenamef  ice-9/session.scm
[
	l		l		l		m		#n		'm		<	 			<	  g  nameg  apropos-foldg  documentationf ÉFolds PROCEDURE over bindings matching third arg REGEXP.

Result is

  (PROCEDURE MODULE1 NAME1 VALUE1
    (PROCEDURE MODULE2 NAME2 VALUE2
      ...
      (PROCEDURE MODULEn NAMEn VALUEn INIT)))

where INIT is the second arg to `apropos-fold'.

Fourth arg FOLDER is one of

  (apropos-fold-accessible MODULE) ;fold over bindings accessible in MODULE
  apropos-fold-exported		   ;fold over all exported bindings
  apropos-fold-all		   ;fold over all bindings C"R     hx     ]*"  _(  C4L 5$   4L4L4L554L55"  "ÿÿ¡ "ÿÿ             g  data
		q g  modules		q g  modules			e g  data			e g  obj			7 g  handle		!	4 g  first?		&	1  g  filenamef  ice-9/session.scm
							$				$	!		&	$	&		,		;		<		?	!	B	.	G	7	I	.	M	!	N	!	S	+	U	!	W		e	 		q	  g  nameg  rec C   h0   Â   ]4	5O L L Q 4L5 6 º       g  fold-module
		/ g  init		/ g  table				/ g  rec			/  g  filenamef  ice-9/session.scm
									)		/	 		/	   C   h     ] O C        g  
init-thunk
		 g  traverse		 g  extract			  g  filenamef  ice-9/session.scm

 			  g  nameg  make-fold-modulesg  documentationf  âReturn procedure capable of traversing a forest of modules.
The forest traversed is the image of the forest generated by root
modules returned by INIT-THUNK and the generator TRAVERSE.
It is an image under the mapping EXTRACT. CR  h   N   ] L  CF       g  filenamef  ice-9/session.scm
			  		
   C   h      ] O 6       ~       g  module
		  g  filenamef  ice-9/session.scm

		 		  g  nameg  apropos-fold-accessible C#R       h      ] 456   y       g  filenamef  ice-9/session.scm

								 		
  g  nameg  root-modules CR       h   d   ]C   \       g  k
		 g  v		  g  filenamef  ice-9/session.scm
¢	 			   C       h      ]4 56 w       g  mod
		  g  filenamef  ice-9/session.scm
¡
	¢	#	¢	 		  g  nameg  
submodules CR4iiii5$R4iiii5%Ri      h0   ¨   ]4 5$   64 5$  4 56C             g  obj
		+  g  filenamef  ice-9/session.scm
ª
	«			«		«		¬			«		!¬	(	)¬	 			+  g  nameg  source C&REG ¡¢£¤¥¦§¨©ª«¬­®¯°± h  x
  ]I4 5$ H(  " P454>  "  G  4>  "  G  4>  "  G  $  4>  "  G  "  4>  "  G  4`>  "  G  4>  "  G  "  ©	(  4'>  "  G  "  	$  L	$  4	>  "  G  "  4
>  "  G  4	>  "  G  		"ÿÿ4>  "  G  4	>  "  G  4'>  "  G  "  		"ÿÿNK(  " lJ$  4>  "  G  "   454>  "  G  4>  "  G  4>  "  G  $  4>  "  G  "  4>  "  G  4`>  "  G  4>  "  G  "  ©	(  4'>  "  G  "  	$  L	$  4	>  "  G  "  4
>  "  G  4	>  "  G  		"ÿÿ4>  "  G  4	>  "  G  4'>  "  G  "  		"ÿÿNK(  " lJ$  4>  "  G  "   454>  "  G  4>  "  G  4>  "  G  $  4>  "  G  "  4>  "  G  4`>  "  G  4>  "  G  "  ©	(  4'>  "  G  "  	$  L	$  4	>  "  G  "  4
>  "  G  4	>  "  G  		"ÿÿ4>  "  G  4	>  "  G  4'>  "  G  "  		"ÿÿNK$  5J$  4>  "  G  "   4>  "  G  K"   $  VJ$  4>  "  G  "   4>  "  G  4>  "  G  4>  "  G  "   "  Ò4 54>  "  G  $  4>  "  G  "  F
$  "  84>  "  G  4>  "  G  4>  "  G  $  "  $  
"  $  4>  "  G  "  4>  "  G  6 p
      g  obj
	 g  t	 g  required-args		-( g  optional-args		-( g  keyword-args		-( g  allow-other-keys?		-( g  rest-arg		-( g  need-punctuation		-( g  len		J g  ls		 Û g  len	À g  ls		Qú g  len	6| g  ls		Çp g  arity	3þ  g  filenamef  ice-9/session.scm
¯
	Ê			Ê		Ê		É		Ì		Í		Î		"Ï	"	(Ð		,Ð		-Ì			?Ò		DÁ		JÁ		MÂ		_Ã		cÃ		hÃ		qÄ		uÓ	:	zÄ	 Å	
 Å	 Æ	
 Æ	 Æ	
 ¡Ç	
 ¥Ç	 ªÇ	
 ³±	 Å²	 Ê²	 Ï²	 Û³	 á´	 âµ	 ú¶	 þ´	»	»	»	¼	¼	¼	½	!½	&½	/¾	4¾	9¾	D¿	J¿	K·	O·	T·	]¸	o¹	³	³	³	Ô	Õ	£Ö	¤Ö	'¨Ö	0­Ö	'ºÁ	ÀÁ	ÃÂ	ÕÃ	ÙÃ	ÞÃ	çÄ	ë×	:ðÄ	üÅ	
 Å	Æ	
Æ	
Æ	
Ç	
Ç	 Ç	
)±	;²	@²	E²	Q³	W´	Xµ	p¶	t´	w»	x»	|»	}¼	¼	¼	½	½	½	¥¾	ª¾	¯¾	º¿	À¿	Á·	Å·	Ê·	Ó¸	å¹	ú³	ý³	³		Ø	Ù	Ú	Ú	'Ú	0#Ú	'0Á	6Á	9Â	KÃ	OÃ	TÃ	]Ä	aÛ	9fÄ	rÅ	
vÅ	wÆ	
{Æ	Æ	
Ç	
Ç	Ç	
±	±²	¶²	»²	Ç³	Í´	Îµ	æ¶	ê´	í»	î»	ò»	ó¼	÷¼	ü¼		½	½	½	¾	 ¾	%¾	0¿	6¿	7·	;·	@·	I¸	[¹	p³	s³	y³	Ü	Ý	Þ	Þ	'Þ	0Þ	'¢ß	¦ß	«ß	¶à	Àá	Æâ	Çâ	'Ëâ	0Ðâ	'Ýã	áã	æã	ïä	å	å	
å	-ç	3ç	6è	;è	@è	Ké	Qé	Rê	Vê	[ê	jë	lë	qé	vì	zì	ì	í	í	í	î	 î	¥î	°ï	¶ï	
¾ð	Àð	Äï	
Çñ	Êñ	Óï	Ôò	
Øò	Ýò	
êó	
îó	óó	
ô	ô	 æ	  g  nameg  arity C'R´µ¶·¸<¹º»¼ÃÆ h  %  ]Q4 5$  Ø$  É$  °$  $  ~	"  6	$  ]	(  R	
45$  4	5"  
45$  4	5"  
 C"ÿÿ"ÿÿ66664 5$  C4 5$   6C            g  proc
	 g  t	 g  w		  ã g  x		  ã g  w		1 Û g  x		1 Û g  w		B Ó g  x		B Ó g  w		S Ë g  x			S Ë g  w	
	t Ã g  t	 ò  g  filenamef  ice-9/session.scm
÷
	þ			þ		þ		ý		ÿ		x			y	 	 	 	+ 	 		 	  	 ¡	 §	+ ©	 °		 Ïÿ	 ì
	 òý	 			
ý			 	  g  nameg  procedure-argumentsg  documentationf  ßReturn an alist describing the arguments that `proc' accepts, or `#f'
if the information cannot be obtained.

The alist keys that are currently defined are `required', `optional',
`keyword', `allow-other-keys?', and `rest'. C(RC      V      g  m
		,  g  filenamef  ice-9/session.scm		
 ½	$	 À	#
´	'
±	0
ç	4
è	8	ë	8
Y	:
W	F
	k	J
« 
x 
\ £
è ö
é#	ô#
&g%
($S
-É[
2±
3É
4t
5¡
5¥	5°¤
5±¨	5Æ§
6»ª
G{¯
Kè÷
  	Kê
   C6 