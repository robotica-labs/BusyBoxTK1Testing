GOOF----LE-4-2.0îA      ] d 4  h      ] g  guile¤	 ¤	g  process-use-modules¤	 ¤	 ¤	g  	aisleriot¤	g  	interface¤	 ¤		 ¤	
g  api¤	
 ¤	 ¤	g  initialize-playing-area¤	g  set-ace-low¤	g  make-standard-double-deck¤	g  shuffle-deck¤	g  add-normal-slot¤	g  DECK¤	g  stock¤	g  waste¤	g  add-blank-slot¤	g  
foundation¤	g  add-carriage-return-slot¤	g  add-extended-slot¤	g  down¤	g  tableau¤	g  
deal-cards¤		
									
								
							
						
					
				
			
		
 -¤	g  map¤	g  flip-top-card¤		
								 	¤	 g  give-status-message¤	!g  new-game¤	"g  set-statusbar-message¤	#g  string-append¤	$g  get-stock-no-string¤	%f     ¤	&g  get-redeals-string¤	'g  _¤	(f  Redeals left:¤	)f   ¤	*g  number->string¤	+g  FLIP-COUNTER¤	,f  Stock left:¤	-g  length¤	.g  	get-cards¤	/g  is-visible?¤	0g  reverse¤	1g  button-pressed¤	2g  move-n-cards!¤	3g  add-to-score!¤	4g  empty-slot?¤	5g  make-visible-top-card¤	6g  complete-transaction¤	7g  	get-value¤	8g  king¤	9g  is-red?¤	:g  get-top-card¤	;g  ace¤	<g  get-suit¤	=g  
droppable?¤	>g  button-released¤	?g  
flip-stock¤	@g  button-clicked¤	Ag  check-to-foundation¤	Bg  autoplay-foundations¤	Cg  button-double-clicked¤	Dg  or-map¤	E	
								 
¤	Fg  delayed-call¤	Gg  game-won¤	Hg  get-hint¤	Ig  game-continuable¤	Jg  strip¤	Kg  	is-black?¤	Lg  
check-plop¤	Mg  	hint-move¤	Ng  	find-card¤	Og  check-a-slot-to-foundations¤	Pg  check-uncover¤	Qg  check-a-foundation-for-uncover¤	Rg  check-foundation-for-uncover¤	Sg  check-empty-tslot¤	Tg  check-move-waste¤	Ug  check-to-foundations¤	Vg  check-simple-foundation¤	Wf  Deal another card¤	Xf  Move waste to stock¤	Yg  	dealable?¤	Zg  get-min-happy-foundation¤	[g  min¤	\g  check-foundation-for-waste¤	]f  Try rearranging the cards¤	^g  get-options¤	_g  apply-options¤	`g  timeout¤	ag  set-features¤	bg  droppable-feature¤	cg  
set-lambda¤C 5  hð7  \  ] 4	 >  "  G      h0  l  ] 4>   "  G  4>   "  G  4>   "  G  4>   "  G  4>  "  G  4>  "  G  4	>   "  G  4
>  "  G  4
>  "  G  4
>  "  G  4
>  "  G  4
>  "  G  4
>  "  G  4
>  "  G  4
>  "  G  4>   "  G  4	>   "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4
>  "  G  4>  "  G  4>   "  G  		 C   d      g  filenamef  	jumbo.scm
	
							#			3			C			I			N			W			Z			\			a			j			z			}					 		 		 		 		 		  	 	 £	 	 ¥	 	 ª	 	 ³	!	 ¶	!	 ¸	!	 ½	!	 Æ	"	 É	"	 Ë	"	 Ð	"	 Ù	#	 Ü	#	 Þ	#	 ã	#	 ì	$	 ï	$	 ñ	$	 ö	$	 ÿ	%		%		%			%		&	"	(	2	*	5	*	9	*	>	*	G	+	J	+	N	+	S	+	\	,	_	,	c	,	h	,	q	-	t	-	x	-	}	-		.		.		.		.		/		/	¢	/	§	/	°	0	³	0	·	0	¼	0	Å	1	È	1	Ì	1	Ñ	1	Ú	2	Ý	2	á	2	æ	2	ï	4	ô	4	ù	4		8		8		8		:	,	;	 \	-
  g  nameg  new-game C!R"#$%&   h      ] 445 45 56        g  filenamef  	jumbo.scm
	=
		>			>	(		?	(		@	(		>			>	 		
  g  nameg  give-status-message C R#'()*+  h      ] 45456       g  filenamef  	jumbo.scm
	B
		C				C			C			C	$		D			D	!		D			C	 
		
  g  nameg  get-redeals-string C&R#',)*-. h    «   ] 45444
5556 £       g  filenamef  	jumbo.scm
	F
		G				G			G			G	"		H			H	!		H	)		H	!		H			G	 		
  g  nameg  get-stock-no-string C$R-/0        h@   â   ]
 $  "   $  45"  $  456C  Ú       g  slot-id
		> g  	card-list		> g  t			,  g  filenamef  	jumbo.scm
	J
		K			K			L			L			M		$	M		0	K		3	N		:	N		<	N	 		>	  g  nameg  button-pressed C1R2345       h    )  ]4 >  "  G   $  # 	
$  4	ÿ>  "  G  "   "   $  "	
$  4>  "  G  "   "   4 5$  "   		$  4 >  "  G  "   C  !      g  
start-slot
	  g  	card-list	  g  end-slot		   g  filenamef  	jumbo.scm
	P
		Q			R		 	R		%	S		)	R		*	T		G	U		K	U		P	V		T	U		U	W		n	X		x	X	 	Y	 	X	 	Z	 	 	  g  nameg  complete-transaction C6R47089:-;< 
     h(  õ  ]"  		$  ~ $  C45$  4455"  $  C45$  C44554455&  C44554455CC$  	
$  45$  p $  C45$  45"  $  C45$  C4	54	455$  454455CCC"ÿþã"ÿþß í      g  
start-slot
	' g  	card-list	' g  end-slot		' g  t		:  g  t	 Ó  g  filenamef  	jumbo.scm
	]
		i				^			j			j				k		%	k		&	l		)	l	*	0	l	%	2	l		5	l		:	k		F	m		P	m		S	n	!	V	n	*	^	n	!	_	o	!	b	o	/	i	o	*	k	o	!	o	m		r	p		u	p	%	}	p		~	q	 	q	- 	q	( 	q	 	q	 	p	 	^	 	^	 	^	 	_	  	^		 ¡	`	 ©	`	 ­	`		 ²	a	 ¶	`		 ¹	b	 ¾	b	% À	b	 Ã	b	 Ç	b	 È	c	 Ó	b	 ß	d	 é	d	 ì	e	 ñ	e	$ ó	e	 ô	f	 ÷	f	$ ÿ	f	 	e		d		g	
	g	%	g		h		h	*	h		h		g	 D	'	  g  nameg  
droppable? C=R=6        h    µ   ]4 5$  
 6C   ­       g  
start-slot
		 g  	card-list		 g  end-slot			  g  filenamef  	jumbo.scm
	t
		u			u			v	 			  g  nameg  button-released C>R?  h      ] 
$  
6C     |       g  slot-id
		  g  filenamef  	jumbo.scm
	y
		z		
	z			{	 		  g  nameg  button-clicked C@R47;<:A h   k  ]
	
$  C45$  4 5"  $  "  B45$  "  24 54455$  4 54455"  $  C 6       c      g  card
	  g  f-slot	  g  t		'	y  g  filenamef  	jumbo.scm
	}
		~				~		 		 		 		" 		' 			5 		? 		E 		L 		O 		W 		X 		\ 		] 		d 		g 	%	o 		p 		q 		}	~	  	&  	 	 	  g  nameg  check-to-foundation CAR4A:53B        h   w  ]	 $  "   		$  h4 5$  C44 5	5$  G4 44 5	5 5$  )4 5$  "  4 5$  6CCC $  6 C   o      g  slot-id
	  g  t		 g  t	a	x  g  filenamef  	jumbo.scm
 
	 		 			 		 		  		* 			- 		0 	#	: 		> 			? 		D 	(	G 	=	Q 	(	T 	"	V 		Z 			[ 		a 		o 		| 		  	  		  	  		 	   g  nameg  button-double-clicked CCRDCEF h      ] 45$  L 6C        g  filenamef  	jumbo.scm
 		 			 	&	 		 		 	 		
  g  nameg  autoplay-foundations-tail CDCEF     h8   Ö   ]O   Q  45$  45$   6CC     Î       g  autoplay-foundations-tail
	
	3  g  filenamef  	jumbo.scm
 
	 		 	$	 		 		 		# 	&	% 		) 		/ 	 		3
  g  nameg  autoplay-foundations CBR GH     h(   }   ] 4>   "  G  45 $  C6        u       g  filenamef  	jumbo.scm
  
	 ¡		 ¢		 ¢		! £	 		!
  g  nameg  game-continuable CIR4  h     ] 4
5$  z45$  n4	
5$  a4	5$  T4	5$  G4	5$  :4	5$  -4	5$   4	5$  4	5$  	6CCCCCCCCCC      g  filenamef  	jumbo.scm
 ¥
	 ¦		 ¦		 §		 ¦		 ¨		! ¦		" ©		, ¦		- ª		7 ¦		8 «		B ¦		C ¬		M ¦		N ­		X ¦		Y ®		c ¦		d ¯		n ¦		t °	 	 
  g  nameg  game-won CGR/J      h    ¥   ]4 5$   6 C             g  	card-list
		  g  filenamef  	jumbo.scm
 ²
	 ³		 ³		 ³		 ³		 µ		 µ		 ´	 			  g  nameg  strip CJR49K:7L8      h   U  ]	$  C45$  "  14 54455&  4 54455"  $  C"  	 645$  4 5$  C"ÿÿ×"ÿÿÓ  M      g  card
	  g  t-slot	   g  filenamef  	jumbo.scm
 ·
	 ¸			 ¸		 º		 º			 »		% ¼		( ¼		0 ¼		4 º			5 ½		< ½		= ¾		@ ¾		H ¾		I ½		R ¸		` Ä		b Ä		b ¸		c À		m ¸		n Á		w Á		{ À		 	 	  g  nameg  
check-plop CLR4/0.LJMN-OP      hØ     ] 	$  C4 5$  "  .444 555$  "  444 55	
5$  , 4 44 555444 55	
564 5$  "  74	4 55$  $44 55$  "  	4
 	5"  $  
 	6 6   
      g  t-slot
	 Õ  g  filenamef  	jumbo.scm
 Æ
	 Ç			 Ç		 É		 É			 Ê		! Ê	%	$ Ê	.	, Ê	%	- Ê	 	/ Ê		3 É			9 Ë		< Ë		? Ë	!	G Ë		K Ë		O Ç		T Ì		Y Ì	-	\ Ì	4	d Ì	-	f Ì		g Ì	I	j Ì	U	m Ì	\	u Ì	U	y Ì	I	{ Ì			| Í	  Í		  Î	  Î	  Î	  Î	  Í		  Ï	 ¡ Ï	& ¨ Ï	  « Ï	 ¯ Í		 µ Ð	 Æ Ç	 Î Ñ		 Ó Ò	 Õ Ò	 /	 Õ  g  nameg  check-uncover CPR49K:7LMQ 	      h   }  ]	
$  C45$  "  H4 54455&  04 54455$  445	
5"  "  $  445	
56 6 u      g  card
	  g  f-slot	   g  filenamef  	jumbo.scm
 Ô
	 Õ			 Õ		 ×		 ×			 Ø		% Ù		( Ù		0 Ù		4 ×			5 Ú		< Ú		= Û		@ Û		H Û		I Ú		M ×			N Ü		Q Ü		[ Ü		i Õ		o Ý		r Ý	)	| Ý		~ Ý		  Þ	1  Þ	 	 	  g  nameg  check-a-foundation-for-uncover CQR4/0.QJR      hh   I  ] 	$  C4 5$  "  .444 555$  "  444 55	5$  44 55	6 6A      g  t-slot
		h  g  filenamef  	jumbo.scm
 à
	 á			 á		 ã		 ã			 ä		! ä	%	$ ä	.	, ä	%	- ä	 	/ ä		3 ã			9 å		< å	.	? å	5	G å	.	K å		O á		R æ	)	U æ	0	] æ	)	a æ			f ç	*	h ç	 		h  g  nameg  check-foundation-for-uncover CRR4/0.LSM- 	       h¨   ì  ] 	$  C4 5$  "  /444 555$  444 55	
5"  $  J4444 55	
55$   6 44 55444 55	
56 6       ä      g  t-slot
	 ¡  g  filenamef  	jumbo.scm
 é
	 ê			 ê		 ì		 ì			 í		! í	 	$ í	)	, í	 	- í		/ í		3 ì			4 î		7 î		: î	(	B î		C î		G î		P ê		Q ï		T ï		W ï	+	Z ï	4	b ï	+	c ï	&	g ï		i ï		m ï			r ð	 	t ð		y ñ		| ñ	'  ñ	  ñ	;  ñ	L  ñ	U  ñ	L  ñ	G  ñ	;  ñ	  ò	 ¡ ò	 +	 ¡  g  nameg  check-empty-tslot CSR49:K7MOT8 
  hÈ   Ø  ]	 	$  "  45$  C4 5$  "  9445544 55&  445544 55"  $   6"  4	5$  	6 64 5$  4455	$   6"ÿÿ½"ÿÿ¹     Ð      g  t-slot
	 Ã g  t		  g  filenamef  	jumbo.scm
 ô
	 õ		 õ			 ö		! õ		$ ø		. ø			4 ù		7 ù		> ù		? ú		B ú		J ú		N ø			O û		R û	!	Y û		Z û		[ ü		^ ü		f ü		g û		p õ		x ý			}		  õ	 		 	 	  õ	  þ	 ¡ õ	 ¢ ÿ	 ¥ ÿ	 ¬ ÿ	 ¯ ÿ	 ³ þ		 » 		 '	 Ã  g  nameg  check-move-waste CTR4UO   hP     ] 	$  C4 5$  	 6 	$  	
	64 5$   6 6ú       g  slot
		P g  f-slot		P  g  filenamef  	jumbo.scm

														!				&
			*		2			3			?		G			L	"	P	 		P	  g  nameg  check-to-foundations CUR47:;<MO h     ]	
$  C45$  44 55"  <44 554455$  44 554455"  $  	 6 6         g  slot
		} g  f-slot		}  g  filenamef  	jumbo.scm

															$		'		,		/		7		8		;		C		D		H		I		L	 	T		U		X	%	`		a		b		k		t			{	.	}	 		}	  g  nameg  check-a-slot-to-foundations CORV47:O        hh   "  ] 	$  C 	$  	
64 5$  "  !44 55$  4 	5"  $   	6 6          g  slot
		d g  happynum		d  g  filenamef  	jumbo.scm

											 			!		)!			/"		2"		:"		="		A!			B#		S		[$			`%	%	d%	 		d	  g  nameg  check-simple-foundation CVR4'W+X     hP   é   ]4
5$  "  
45   $   C45$  C$  
45 CC       á       g  t
		I  g  filenamef  	jumbo.scm
'
	(		(		)		)		)		)		(		(*		1*		7+		;*		=,		A,		C,		F,	 		I
  g  nameg  	dealable? CYR4Z[7:      h@   ô   ] 	
$  C4 5$  	C 4	44 5556       ì       g  fslot
		9 g  value		9  g  filenamef  	jumbo.scm
.
	/			/		1			/		!3	&	"4	&	)4	6	,4	A	44	6	54	1	74	&	93	 		9	  g  nameg  get-min-happy-foundation CZR49:K7LM\ 	    h¨   ¯  ]	 	
$  "  45$  C4 5$  "  P44 554455&  444 554455$  44 5	
5"  "  $   44 5	
56 6     §      g  f-slot
	 £ g  t		  g  filenamef  	jumbo.scm
6
	7		7			8		!7		$:		.:			4;		7;		?;		@<		C<		J<		N:			O=		R=		Z=		[>		^>	!	e>		f>		g=		k:			l?		o?		y?	 7	 @	 @	) @	 @		 ¡A	( £A	 "	 £  g  nameg  check-foundation-for-waste C\RVZ8PSTR\YU']  h¨     ]44	55  $   C4	
5  $   C4	
5  $   C4	
5  $   C4	
5  $   C4	5  $   C4	5   $   C4
	5  $   C
45 C         g  t
	 ¥ g  t
	# ¥ g  t
	5 ¥ g  t
	G ¥ g  t
	Y ¥ g  t
	k ¥ g  t
	{ ¥ g  t
  ¥  g  filenamef  	jumbo.scm
C
	D		D	!	D		D		E		#D		/F		5D		AG		GD		SH		YD		eI		kD		wJ		{D	 K	 D	 L	 L	 ¡L	 ¤L	 	 ¥
  g  nameg  get-hint CHR  h   T   ] C    L       g  filenamef  	jumbo.scm
N
 		
  g  nameg  get-options C^R       h   l   ]C    d       g  options
		  g  filenamef  	jumbo.scm
Q
 		  g  nameg  apply-options C_R       h   P   ] C    H       g  filenamef  	jumbo.scm
T
 		
  g  nameg  timeout C`R4aibi>  "  G  ci!i1i>i@iCiIiGiHi^i_i`i=i6 T      g  filenamef  	jumbo.scm		
ï	
®	=
}	B
f	F
	¥	J
	P
Ð	]
À	t
o	y
	}
º 
¡ 
`  

 ¥
è ²
è ·
ý Æ
"( Ô
#ü à
&· é
){ ô
*å
-

.µ
0'
1_.
3Ú6
6DC
6¯N
77Q
7£T
7¤W
7ïY
 #	7ï
   C6 