GOOF----LE-4-2.0_M      ]  4  h³	      ] g  guile¤	 ¤	g  define-module*¤	 ¤	 ¤	g  system¤	g  xref¤	 ¤		g  filenameS¤	
f  system/xref.scm¤	g  importsS¤	g  base¤	g  pmatch¤	 ¤	 ¤	g  compile¤	 ¤	 ¤	g  vm¤	g  program¤	 ¤	 ¤	g  srfi¤	g  srfi-1¤	 ¤	 ¤	 ¤	g  exportsS¤	g  *xref-ignored-modules*¤	g  procedure-callees¤	g  procedure-callers¤	 g  source-closures¤	!g  source-procedures¤	" ! ¤	#g  set-current-module¤	$# ¤	%# ¤	&g  memq¤	'g  program-objects¤	(g  vector-length¤	)g  make-vector¤	*g  	decompile¤	+g  program-objcode¤	,g  toS¤	-g  assembly¤	.g  load-program¤	/g  for-each¤	0g  toplevel-set¤	1g  toplevel-ref¤	2g  program?¤	3g  fold¤	4g  program-callee-rev-vars¤	5g  	variable?¤	6g  module-variable¤	7g  program-module¤	8g  the-root-module¤	9g  nested-ref-module¤	:g  resolve-module¤	;g  module-public-interface¤	<g  procedure-callee-rev-vars¤	=g  *callers-db*¤	>g  make-hash-table¤	?g  *module-callees-db*¤	@g  *tainted-modules*¤	Ag  value-history¤	BA ¤	CB ¤	Dg  module-name¤	Eg  member¤	Fg  on-module-modified¤	Gg  	hashq-ref¤	Hg  assoc¤	Ig  
hashq-set!¤	Jg  
add-caller¤	Kg  assoc-remove!¤	Lg  forget-callers¤	Mg  	hash-set!¤	Ng  append¤	Og  hash-ref¤	Pg  add-callees¤	Qg  ensure-callers-db¤	Rg  untaint-modules¤	Sg  hash-for-each¤	Tg  module-observers¤	Ug  module-observe¤	Vg  module-for-each¤	Wg  
procedure?¤	Xg  filter¤	Yg  variable-bound?¤	Zg  module-submodules¤	[g  current-module¤	\g  error¤	]f  /expected a variable, symbol, or (modname . sym)¤	^g  program-sources¤	_g  procedure-sources¤	`g  *closure-sources-db*¤	ag  *sources-db*¤	bg  *module-sources-db*¤	cg  *tainted-sources*¤	dg  on-source-modified¤	eg  
hashv-set!¤	fg  	hashv-ref¤	gg  
add-source¤	hg  delq¤	ig  hashv-remove!¤	jg  forget-source¤	kf  unexpected source format¤	lg  add-sources¤	mg  and=>¤	ng  vector->list¤	og  hashq-remove!¤	pg  forget-sources¤	qg  ensure-sources-db¤	rg  untaint-sources¤	sg  sort!¤	tg  hash-map->list¤	ug  cons¤	vg  lines->ranges¤	wg  reverse¤	xg  <=¤	yg  lookup-source-procedures¤	zg  canonicalizationS¤	{z	¤	|{ ¤	}g  relative¤	~g   %file-port-name-canonicalization¤	g  catch¤  ¤  ¤ g  open-input-file¤ g  port-filename¤C 5hA    ]4	
"5 4% >  "  G   &     h      ]4 5$  C C       g  x
		 g  y		  g  filenamef  system/xref.scm
	"			#			#			#	 			  g  nameg  	cons-uniq C'()*+,-./01       h   ÿ   ]!"  @ $  7  &  #$  (  L ¤CCCCC $  =  &  '$  (  L ¤C"ÿÿ"ÿÿ"ÿÿ|"ÿÿx  ÷       g  x
	  g  vx		D g  vy			D g  vx		,	@ g  vy		,	@ g  vx	S  g  vy		S  g  vx		l  g  vy		l   	g  filenamef  system/xref.scm
	-			.		<	0	'	F	.		|	/	' 	.	 	    C2345&6789:;        h@  ¡  ]4 5$ )45445544 55$  g	&  K$  @$  )	4
O 	>  "  G  	"   "   "   "   " $  C4£5$  44£55"ÿÿÈ£$ 5£45$  #45$  "  "ÿÿ$  R44 5		$  	"  	5		$  4	5$  "  	"  "ÿÿ1$  
	
$  

$  x(  f445	5$  4$  45"  5"  $  45$  "  "  "ÿþCCCC"ÿþ}
"ÿþsC      g  prog
	@ g  	cons-uniq	@ g  t		@ g  n		9> g  progv		9> g  asm		9> g  vx		L © g  vy		L © g  vy		b  g  vy			r  g  i	 ±4 g  out	 ±4 g  obj	 ÷' g  t		8K g  v		O g  vx		% g  vy	
% g  vx	# g  vy	# g  vx	¯! g  vy	¯! g  m	Ç g  v	î  g  filenamef  system/xref.scm
	!
		%			$			'			(		 	(	$	)	(		*	)		-	)	 	7	)	<	9	)		9	'	
	B	*		y	,	 ±	2	 ¶	4	 º	3	 ¾	5	 Å	5	 Ç	5	 Ë	3	 Î	6	 Ï	7	 Ö	8	 Ý	8	3 ß	8	 á	7	 é	6	 î	9	 ò	3	 ÷	:	 ÷	:	 ú	;		;		<		#		#		#	'	<	*	?	%.	>	/	@	"2	@	78	@	3O	@	"O	@	T	C	Z	C	&[	#	g	#	r	#		C		>	º	F	#½	F	6À	F	FÃ	F	6Ç	F	#Ç	F	Ï	H	#Ð	I	(Ø	J	)Ù	K	-é	I	(î	F	ó	N	ù	O	ú	#		#		#		N	*	P	4	P	4	2	6	2	 >	2	?	Q		 N	@  g  nameg  program-callee-rev-vars C4R24    h      ]4 5$   6C          g  proc
		  g  filenamef  system/xref.scm
	S
		U			T			U			V		 		  g  nameg  procedure-callee-rev-vars C<R<     hH   Y  ]"  0(  C$  "ÿÿÝ"ÿÿÐ4 5"ÿÿÀ  Q      g  prog
		F g  in		6 g  out			6  g  filenamef  system/xref.scm
	X
		Z			[			\			\			[			]			]	,		]		!	]		)	]		,	^		6	^		6	Z		7	Z		>	Z	6	F	Z	 		F  g  nameg  procedure-calleesg  documentationf  1Evaluates to a list of the given program callees. CR=R4>i5 ?R@RCRDE@ hH   Ñ   ]	4 545$  "  45$  "  $  	 CC       É       g  m
		A g  name			A  g  filenamef  system/xref.scm
	h
		i				i			j			j			k		*	j		2	l		6	j		;	m	 	=	m	 		A  g  nameg  on-module-modified CFRG=HEI      hh   E  ]4 5$  =45$  45$  CC  6   6      =      g  callee
		b g  caller		b g  mod-name			b g  all-callers			b g  callers			P  g  filenamef  system/xref.scm
	o
		p			p			q			s			s		$	t	
	%	u		1	u		:	v	1	;	v	$	<	v		K	x	 	N	x		P	w		]	r	(	b	r	 		b	  g  nameg  
add-caller CJRI=KG    h    ·   ] 44 556     ¯       g  callee
		 g  mod-name		  g  filenamef  system/xref.scm
	z
			|			|			|	;		|			|			{	 			  g  nameg  forget-callers CLRM?NO  h    º   ]4 4556     ²       g  callees
		 g  mod-name		  g  filenamef  system/xref.scm
	~
		 		 		 	D	 		 				 			  g  nameg  add-callees CPRQ//L     h   a   ] L 6      Y       g  callee
		
  g  filenamef  system/xref.scm
 		
 	 		
   CO?Q      h0      ]4 O 4 5>  "  G   6              g  m
		)  g  filenamef  system/xref.scm
 		 		 		 	.	 		 		) 	 		)  g  nameg  untaint C@      h0      ] 4>  "  G  4>  "  G   C    z       g  filenamef  system/xref.scm
 
	 		 		( 		* 	 		,
  g  nameg  untaint-modules CRR:SI&FTUDVGWIXY</J h   a   ] LL 6    Y       g  callee
		  g  filenamef  system/xref.scm
 		 	 		   CP 
 h   9  ]$  M $  4M 5"  $  C45$  RM $  4M >  "  G  "   44554LO >  "  G  	L6CC  1      g  sym
	  g  var	  g  x		  g  t		"  g  callees		e   g  filenamef  system/xref.scm
 ¡		 	
		 		 		 	
	 		 		" 		. 		8 		> 		? 		X 		] 	%	e 		e 		h 	  	 	 	   CG 
 h    ü   ]
"  vM$  4M>  "  G  "   4455$  "  4>  "  G  454LO >  "  G  L 6M$  4	M5$  C"ÿÿq"ÿÿm       ô       g  name
	  g  sub	  g  name		V	v  g  filenamef  system/xref.scm
 §		 		 		& 		+ 	(	3 		7 		< 	
	P  		V  		Y ¡		| «		| ¨		  ¨	  ¨	 	 	   CZ       h      ]L LO 4 56  |       g  mod
		  g  filenamef  system/xref.scm
 ¥		 ¬		 ¦	 		  g  nameg  visit-submodules C=>I&FTUDVGWIXY</J     h   a   ] LL 6    Y       g  callee
		  g  filenamef  system/xref.scm
 		 	 		   CP 
 h   9  ]$  M $  4M 5"  $  C45$  RM $  4M >  "  G  "   44554LO >  "  G  	L6CC  1      g  sym
	  g  var	  g  x		  g  t		"  g  callees		e   g  filenamef  system/xref.scm
 ¡		 	
		 		 		 	
	 		 		" 		. 		8 		> 		? 		X 		] 	%	e 		e 		h 	  	 	 	   C    hÈ   ¼  ]! $  4 5"  HO Q  $  "  $  4è5 4è5K456 $  aJ$  4J>  "  G  "   4455$  "  4	>  "  G  4
5O 6C     ´      g  mod-name
	 Ã g  mod	 Ã g  visited		 Ã g  visit-submodules		" Ã g  name	 ° Á  g  filenamef  system/xref.scm
 
	 			 		 		/ ®		7 ®		; ®		< ¯		E ¯		F °		O °		R ±		U ±	-	X ±		Z ±		` ®		f 		g 	  	  	(  	  	  	
 ª  	 °  	 Á ¡	 	 Ã  g  nameg  ensure-callers-db CQR56[:\]RG= 
  hx   N  ]4 5$   "  I $  445  5"  2 $  "  4455"  	4 54>   "  G  	6 F      g  var
		w g  vx	7	O g  vy		7	O g  v	\	w  g  filenamef  system/xref.scm
 ´
	 ¹		 ¹		 º		 ¹		 º	 	 º	1	& º	 	- ¼		< ¾		? ¾	&	I ¾		T À		X À		\ À		\ ¹		_ Á		u Â		w Â	 		w  g  nameg  procedure-callersg  documentationf  ÷Returns an association list, keyed by module name, of known callers
of the given procedure. The latter can specified directly as a
variable, a symbol (which gets resolved in the current module) or a
pair of the form (module-name . variable-name),  CR2^       h      ]4 5$   6C          g  proc
		  g  filenamef  system/xref.scm
 Í
	 Ï		 Î		 Ï		 Ð		 		  g  nameg  procedure-sources C_R`RaR4>i5 bRcRDEc    hH   Ü   ]	4 545$  "  45$  "  $  	 CC       Ô       g  m
		A g  name			A  g  filenamef  system/xref.scm
 Û
	 Ü			 Ü		 Ý		 Ý		 Þ		* Ý		2 ß		6 Ý		; à	 	= à	 		A  g  nameg  on-source-modified CdRO>Mef   hX   =  ]45$  "  "45 4>  "  G   456   5      g  proc
		U g  file		U g  line			U g  db			U g  t			> g  table			; g  
file-table		>	U  g  filenamef  system/xref.scm
 â
	 ã		 ã		 ä	%	 ä		! å		> ã		I é		P é	6	R é		S é		U ç	 		U	  g  nameg  
add-source CgROhfei  hH   C  ]45$  ,4 455$  
66C       ;      g  proc
		A g  file		A g  line			A g  db			A g  
file-table			A g  procs		$	?  g  filenamef  system/xref.scm
 ë
	 ì		 ì		 í		 î		 î	 	  î	;	" î	 	$ î		$ î		) ï		- ï	
	7 ð		? ñ	 		A	  g  nameg  forget-source CjR_IOb>M/g\k hX   ß   ]! $  F  $  .$  L L6 6 6 6 ×       g  source
		W g  vy		O g  vx			G g  vy			G g  vx		-	?  g  filenamef  system/xref.scm
 ÿ		 		?		C	%	G		K	%	O		S	%	W	 
		W   CWl`       h    p   ]4 5$  
 L 6C       h       g  obj
		  g  filenamef  system/xref.scm
							 		   C2m'n     hÀ     ]4 5$  o445$  "  "45 4>  "  G   >  "  G  4 O >  "  G  "   	O 4
 5$  44 55"  $  "  6          g  proc
	 ¼ g  mod-name	 ¼ g  db		 ¼ g  sources			 ¼ g  t			Q g  table		1	N g  t	 ¨ º  g  filenamef  system/xref.scm
 ó
	 ô			 ô		 õ		 õ		 ø	
	 ø		 ø		- ù	'	1 ù		4 ú		Z ø	
	c ÿ	
 		 		 
	 
	 £
	 ¨		 ·	 ¼	 	 ¼	  g  nameg  add-sources ClROb/j\k        hX   ß   ]! $  F  $  .$  L L6 6 6 6 ×       g  source
		W g  vy		O g  vx			G g  vy			G g  vx		-	?  g  filenamef  system/xref.scm
				?		C	%	G		K	%	O		S	%	W	 
		W   CGoWp`   h    p   ]4 5$  
 L 6C       h       g  obj
		  g  filenamef  system/xref.scm
							 		   C2m'n     h   ]  ]45$  r4 O 4 5>  "  G  4 >  "  G  O 4 5$  4	4
 55"  $  "  6C U      g  proc
	  g  mod-name	  g  db		  g  	mod-table		  g  t		q   g  filenamef  system/xref.scm

								
	 		'	.	)		.	
	7	
	T		^		_		b		l		q	  	 	
 	 	  g  nameg  forget-sources CpRq//pa        h   _   ] L 6    W       g  proc
		  g  filenamef  system/xref.scm
$		$	 		   CObS h   k   ] C   c       g  proc
		 g  sources		  g  filenamef  system/xref.scm
(	$ 			   Cq        hH   ´   ]	4 O 4 5$  45"  >  "  G   6       ¬       g  m
		A g  t		.  g  filenamef  system/xref.scm
#		$		&		%		(		+)		3$		A*	 			A  g  nameg  untaint Cc       h0      ] 4>  "  G  4>  "  G   C    z       g  filenamef  system/xref.scm
"
	+		,		(-		*-	 		,
  g  nameg  untaint-sources CrR&dTUDVWla   h(   ´   ]
$  45$  
L 6CC ¬       g  sym
		' g  var		' g  x			%  g  filenamef  system/xref.scm
5		6			6			7		7		8		8		#9	 			'	   C hH   Ç   ]	44 55$  "  4 >  "  G  4 5O  6      ¿       g  mod
		B g  name	3	B  g  filenamef  system/xref.scm
0		1		1	&	1		1		2		-3		33		B4	 
		B  g  nameg  visit-module C>SGI    h@      ]4M5$  C4M>  "  G  4L >  "  G  M6         g  name
		> g  sub		>  g  filenamef  system/xref.scm
B			C		C		E		'F		>G	 		>	   CZ       h0      ]M$  "  45 NLL LO 4 56       z       g  mod
		)  g  filenamef  system/xref.scm
>		?		@		@		!H			)A	 		)   Ca`>:   hp   J  ]HHO K $  "  $  "  $  4è5 4è5 J456 $  4 56C     B      g  mod-name
		k g  visit-submodules		k g  visit-module				k g  visited		
	  g  filenamef  system/xref.scm
/
	
=		"J			5J	0	9J		:K	$	CK			DL		ML			PM		SM	+	VM		XM			^J		aN	 	iN	 		k  g  nameg  ensure-sources-db CqR>S/GI   hH   Ð   ]	4L  5$  &L$  	L"   L$  LCCL  LL6  È       g  proc
		F g  t		F  g  filenamef  system/xref.scm
U		W		V	
	Y		Y		Y		"Z		+[		,[		0[		5\		D^	$	F^	 		F   C     h   r   ]L  O 6       j       g  line
		 g  procs		  g  filenamef  system/xref.scm
S		T	 			   Cstu      h   }   ] C    u       g  x
		 g  y		  g  filenamef  system/xref.scm
b		b			b	%	b	 			   C        h8   µ   ]	45 4O  >  "  G  456       ­       g  
file-table
		1 g  ranges		1  g  filenamef  system/xref.scm
P
	Q		Q		
R		%a		1a	 		1  g  nameg  lines->ranges CvROwxv    hp   §  ]4 5"  @(  645$  "ÿÿÍ"ÿÿÀ$  45"  "ÿÿ¥          g  
canon-file
		l g  line		l g  db			l g  
file-table			l g  ranges			Q g  procs			Q  g  filenamef  system/xref.scm
d
	e		e		f		h		i		j		#j		*j	 	.j		2h		5k		8k		<k		Dk		Gm		Qm		Qf		Wf		Xf	$	cf	?	dg		lf	 		l	  g  nameg  lookup-source-procedures CyR|}q~     h   L   ] L 6D       g  filenamef  system/xref.scm
r		r	% 		
   C  h   X   -  1  3 C     P       g  args
			  g  filenamef  system/xref.scm
r	 			


   Cy`       hh   +  - /   0   3 #  4>  "  G  Y4 O 5Z$  45"   	
6   #      g  file
		e g  line		e g  canonicalization			e g  port		D	e g  file		Y	e  g  filenamef  system/xref.scm
o
	o	<	p		5r		Dq		Ls		Ms		Yq		et	 
		e	
g  canonicalizationS	   g  nameg  source-closures C R|}q~ h   L   ] L 6D       g  filenamef  system/xref.scm
y		y	% 		
   C  h   X   -  1  3 C     P       g  args
			  g  filenamef  system/xref.scm
y	 			


   Cya       hh   -  - /   0   3 #  4>  "  G  Y4 O 5Z$  45"   	
6   %      g  file
		e g  line		e g  canonicalization			e g  port		D	e g  file		Y	e  g  filenamef  system/xref.scm
v
	v	>	w		5y		Dx		Lz		Mz		Yx		e{	 
		e	
g  canonicalizationS	   g  nameg  source-procedures C!RC      }      g  m
		,  g  filenamef  system/xref.scm		
	¬	!

u	S
,	X
0	a
1	c	:	c
;	e	>	e
@	g	C	g
t	h
@	o
2	z
%	~
] 
G 
 1 ´
 ú Í
 þ Ó
! Õ
! ×	! ×
! Ù	! Ù
"O Û
$  â
%¦ ë
* ó
. 
1"
6­/
: P
<Rd
>æo
Axv
 %	Az
   C6 