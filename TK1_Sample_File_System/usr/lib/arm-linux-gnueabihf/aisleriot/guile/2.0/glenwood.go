GOOF----LE-4-2.02H      ] q 4     h¸      ] g  guile¤	 ¤	g  process-use-modules¤	 ¤	 ¤	g  	aisleriot¤	g  	interface¤	 ¤		 ¤	
g  api¤	
 ¤	 ¤	g  BASE-VAL¤	g  call-with-deferred-observers¤	 ¤	 ¤	g  module-export!¤	 ¤	 ¤	g  current-module¤	 ¤	 ¤	 ¤	g  variable-list¤					 ¤	g  
foundation¤						
 ¤	g  tableau¤	g  initialize-playing-area¤	g  set-ace-low¤	g  make-standard-deck¤	 g  shuffle-deck¤	!g  add-normal-slot¤	"g  DECK¤	#g  stock¤	$g  waste¤	%g  add-blank-slot¤	&g  add-carriage-return-slot¤	'g  add-extended-slot¤	(g  right¤	)g  reserve¤	*g  down¤	+g  deal-cards-face-up¤	,																	
 ¤	-g  give-status-message¤	.g  new-game¤	/g  set-statusbar-message¤	0g  string-append¤	1g  get-stock-no-string¤	2f     ¤	3g  get-redeals-string¤	4g  get-base-string¤	5g  _¤	6f  Stock left:¤	7f   ¤	8g  number->string¤	9g  length¤	:g  	get-cards¤	;f  Redeals left:¤	<g  FLIP-COUNTER¤	=f  Base Card: Ace¤	>f  Base Card: Jack¤	?f  Base Card: Queen¤	@f  Base Card: King¤	Af   ¤	Bf  Base Card: ¤	Cg  empty-slot?¤	Dg  is-visible?¤	Eg  button-pressed¤	Fg  	get-value¤	Gg  get-suit¤	Hg  get-top-card¤	Ig  ace¤	Jg  king¤	Kg  	is-black?¤	Lg  is-red?¤	Mg  reverse¤	Ng  
droppable?¤	Og  move-n-cards!¤	Pg  add-to-score!¤	Qg  button-released¤	Rg  
flip-stock¤	Sg  button-clicked¤	Tg  
deal-cards¤	Ug  move-to-foundation¤	Vg  	place-ace¤	Wg  place-found¤	Xg  button-double-clicked¤	Yg  game-won¤	Zg  get-hint¤	[g  game-continuable¤	\f  Move waste back to stock¤	]f  Deal a new card from the deck¤	^g  	dealable?¤	_f  8Select a card from the reserve for first foundation pile¤	`g  base-not-set?¤	ag  check-a-foundation¤	bg  find-foundation¤	cg  to-foundations¤	dg  	hint-move¤	eg  find-empty-slot¤	fg  check-a-tableau-with-single¤	gg  check-a-tableau-pile¤	hg  
to-tableau¤	if  8Move a card from the reserve onto the empty tableau slot¤	jg  empty-tableau?¤	kg  get-options¤	lg  apply-options¤	mg  timeout¤	ng  set-features¤	og  droppable-feature¤	pg  
set-lambda¤C 5   hp=  p  ] 4	 >  "  G  
R4      h   >   ] 45 6   6       g  filenamef  glenwood.scm
	
 		
   C>  "  G  iiRR !"#$%&'()*+,-      hø  Ä  ] 4>   "  G  4>   "  G  4>   "  G  4>   "  G  4>  "  G  4>  "  G  4	>   "  G  4
>  "  G  4
>  "  G  4
>  "  G  4
>  "  G  4>   "  G  4>  "  G  4	>   "  G  4	>   "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>   "  G  4>  "  G  4>   "  G  4>  "  G  4>   "  G  4>  "  G  4
>  "  G  4>   "  G  		 C¼      g  filenamef  glenwood.scm
	
							#			3			C			I			N			W			Z			\			a			j	 		z	"		}	"			"	 	"	 	#	 	#	 	#	 	#	  	$	 £	$	 ¥	$	 ª	$	 ³	%	 ¶	%	 ¸	%	 ½	%	 Æ	'	 Ö	)	 Ù	)	 Ý	)	 â	)	 ë	+	 û	,		.		.		.		.	 	/	#	/	'	/	,	/	5	0	8	0	<	0	A	0	J	1	M	1	Q	1	V	1	_	3	o	4	r	4	v	4	{	4		5		6		6		6	 	6	©	7	¹	8	¼	8	À	8	Å	8	Î	:	Ó	:	Ø	:	á	<	÷	>	 I	ø
  g  nameg  new-game C.R/01234 h       ] 445 45 45 56         g  filenamef  glenwood.scm
	B
		C			C	(		D	(		E	(		F	(		G	(		C			C	 
		
  g  nameg  give-status-message C-R056789:     h    ®   ] 45444
5556 ¦       g  filenamef  glenwood.scm
	I
		J				J			J			J	"		K			K	!		K	)		K	!		K			J	 		
  g  nameg  get-stock-no-string C1R05;78<       h      ] 45456       g  filenamef  glenwood.scm
	M
		N				N			N			N	$		O			O	!		O			N	 
		
  g  nameg  get-redeals-string C3R5=>?@A0B8        hp   2  ] "  >$  6	$  6	$  6	$  6C$  	$  4	54
56"ÿÿ"ÿÿ *      g  filenamef  glenwood.scm
	Q
	
	U				R			V			V				W				R		!	X		#	X			(	Y			,	R		0	Z		2	Z			7	[			;	R		?	\		A	\			C	]		D	R		H	R		L	R		Q	S		U	R			X	T		\	T		^	T		_	T	)	g	T		 		o
  g  nameg  get-base-string C4RCD9:       hH  ­  ]4 5$  C 	$  "  / 	$  "   	$  "   	$  C4	5$  "  45$  45$  "  d4544 55$  K 	$  "  1 	$  "   		$  "   	
"  "  $  C45$  6 	$  C 	$  C 	$  C 	CC¥      g  slot-id
	H g  	card-list	H g  t			S g  t		&	P g  t		8	M g  t		~ ó g  t	 © ë g  t	 ¼ è g  t	 Ï å g  t	 øH g  t	F g  t	%F g  t	5F  g  filenamef  glenwood.scm
	_
		`			`			a			a		&	b		&	a		8	c		8	a		J	d		W	`		Z	e		d	e		j	f		o	f		q	f		u	e		v	g		~	g		~	g	 	h	 	i	 	i	$ 	i	  	h	 ¤	h	 ©	j	 ©	j	 ¼	k	 ¼	j	 Ï	l	 Ï	j	 â	m	 ø	e		n		n		n		o		o	%	p	%	o	5	q	5	o	E	r	 -	H	  g  nameg  button-pressed CER9CFGHIJKLM      h8  ^  ]+ $  C45$ x45$  		"  $  " N4	5$  " >45$ \	$  "  	$  "  	$  45"  $  " 4 5$  "  H 	$  "  1 	$  "   		$  "   	
$    $  "  74	5$  +4	5$  4	5$  4	5"  "  "  $  K	$  "  1	$  "  		$  "  	
"  "  "  $  "  Å45$  "  µ	$  "  /	$  "  	$  "  	$  l445545$  N454455$  "  #45$  4455"  "  "  "  $  C4 5$  "  X $  "  D 	$  "  / 	$  "   	$  "   	$ %	$  "  /	$  "  		$  "  	
$  Û45$  "  v4	4554
455&  X44554455$  "  (4455$  4455"  "  $  C45$  = $  C4	5$   4	5$  4	5$  	6CCCCCCV      g  
start-slot
	8 g  	card-list	8 g  end-slot		8 g  t		/ g  t		]  g  t		o  g  t	 ° g  t	 ° g  t	 Â g  t	 Õ g  t	 è þ g  t	X g  t	a£ g  t	t  g  t	 g  t	µ g  t	Ø g  t	ê g  t	ü g  t	J~ g  t	8 g  t	¥ g  t	¶
 g  t	È g  t	Ú g  t	ì g  t	U g  t	(R g  t	:O g  t	¡Ú g  t	ß4 g  t	ú2   g  filenamef  glenwood.scm
	t
		u			u			v			v			v			w		%	w		*	x		/	w		>	y		H	y		N	z	"	X	z		]	{	/	]	{	+	o	|	/	o	{	+ 	}	/ 	{	& 	~	7 	~	B 	~	7 	~	+ 	{	" ª		/ °		+ Â 	8 Â 	4 Õ 	8 Õ 	4 è 	8 è 	4 û 	8 	/		& 	4 	/ 	+ 	4) 	/* 	44 	/5 	4? 	/@ 	4\		&a 	/a 	+t 	/t 	+ 	/ 	+ 	/µ	z	Ä 	'Î 	Ø 	&Ø 	"ê 	&ê 	"ü 	&ü 	" 	& 	 	% 	/' 	%( 	%- 	// 	%0 	"4 	5 	): 	4< 	)= 	.@ 	9H 	.I 	)J 	&J 	"Y 	.^ 	9` 	.c 	+g 	&h 	.k 	9s 	.v 	+	v	 	¥ 	¶ 	¶ 	È 	È 	Ú 	Ú 	ì 	ì 	þ 	 	  	  	( ¡	(  	: ¢	:  	L £	Y 	Z ¤	d ¤	j ¥	m ¥	)u ¥	v ¦	y ¦	, ¦	' ¦	 ¤	 §	% §	5 §	0 §	% §	  ¨	  ¨	+  ¨	 ¡ §	¡ §	° ©	%³ ©	5º ©	0¼ ©	%¿ ©	"Ã ©	Ä «	%Ç «	0Ï «	%Ò «	"ß ¤	ë ­	õ ­	ù ®	"ú ®	ú ®	 ¯	" ¯	 °	" ¯	 ±	"& ¯	, ²	" 	8	  g  nameg  
droppable? CNRNCFOP       h     ]4 5$  u	$  "  #45$  "  45 $  64 5$  %	$  C	$  C6CCC      g  
start-slot
	  g  	card-list	  g  end-slot		  g  t			J g  t		,	G g  t		b  g  t		r   g  filenamef  glenwood.scm
 ´
	 µ		 µ		 ¶		 ¶		 ¶		% ·		, ·		, ¶		: ¸		? ¸	%	A ¸		C ¸		N µ		O ¹		] µ		b º		b º		r »		r º	  ¼	 	 	  g  nameg  button-released CQRCR     h       ]4	5$  C 
$  
6C       g  slot-id
		   g  filenamef  glenwood.scm
 ¾
	 ¿		 ¿		 À		 ¿		 Á	 		   g  nameg  button-clicked CSRTP        h    È   ]4  >  "  G  6  À       g  
start-slot
		 g  	card-list		 g  end-slot			  g  filenamef  glenwood.scm
 Ã
	 Ä		 Ä		 Ä		 Å	 			  g  nameg  move-to-foundation CURCU     hX   ø   ]4	5$    	64	5$    	64	5$    	6  	6 ð       g  card
		W g  slot		W  g  filenamef  glenwood.scm
 Ç
	 È		 È		 É		 É		 Ê	
	% Ê		. Ë	#	2 Ë	
	3 Ì		= Ì	
	F Í	'	J Í		S Î	'	W Î	 		W	  g  nameg  	place-ace CVRCFHIJGUW 	h    ©  ]45$  "  c454455$  "  !45&  4455"  $  454455"  $    6	$  C 6    ¡      g  slot
	  g  top-card	  g  search		  g  t		'	X  g  filenamef  glenwood.scm
 Ð
	 Ñ		 Ñ		 Ò		 Ò		 Ó		 Ó		& Ó		' Ò		' Ò		5 Ô		A Ô		B Õ		E Õ	$	M Õ		P Õ		\ Ñ		] Ö		d Ö	$	g Ö	.	o Ö	$	p Ö		y Ñ	  Ø	!  Ø	  Ù	
  Ù	  Û	%  Û	
 	 	  g  nameg  place-found CWRCHFVW    h°     ]! 
$  "  D 	$  "  / 	$  "   	$  "   	$  "  4	5$  "  4 5$  (4 545&   6 	6C             g  slot-id
	 © g  t		Z g  t			W g  t		*	T g  t		<	Q g  top-card  §  g  filenamef  glenwood.scm
 ß
	 à		 à		 á		 à		* â		* à		< ã		< à		N ä		^ à		d å		n à		t æ		{ æ		 à	  ç	  ç	  è	  è	  é	 § ê	 	 ©  g  nameg  button-double-clicked CXR-YZ     h(      ] 4>   "  G  45 $  C6        x       g  filenamef  glenwood.scm
 ï
	 ð		 ñ		 ñ		! ò	 		!
  g  nameg  game-continuable C[R9:     hX   ÿ   ] 	44	55$  :	44	55$  %	44	55$  	44	55CCCC       ÷       g  filenamef  glenwood.scm
 ô
	 õ		 õ		 õ		 õ		 õ		 ö		 ö		# ö		$ ö		( õ		+ ÷		. ÷		6 ÷		7 ÷		; õ		> ø		A ø		I ø		J ø	 		Q
  g  nameg  game-won CYRC<5\]        h@   Ó   ] 4
5$  "$  45$  C
45 CC
45 C      Ë       g  filenamef  glenwood.scm
 ú
	 û		 û		 ý		 ý		 þ		 ý	
	" ÿ		& ÿ		( ÿ		+ ÿ	
	0 ü		4 ü		6 ü		9 ü	 		:
  g  nameg  	dealable? C^RC5_        h       ] 4	5$  
45 CC            g  filenamef  glenwood.scm

												 		
  g  nameg  base-not-set? C`RCaGHFIJ    h   Ï  ]
	$  C45$  	 644 554455&  J44 554455$  C44 55$  4455CC 6Ç      g  slot-id
	  g  foundation-id	  g  t		W   g  filenamef  glenwood.scm

												
	%	!
			"		%		-		.		1		9		=		>		A		I		J		M	 	U		V		W		W			c		f	 	n		q		u		v		y	  	 	 	( 	 #	 	  g  nameg  check-a-foundation CaRCGHb  h8     ]45$  "   4455$  C 6    ý       g  suit
		4 g  foundations		4  g  filenamef  glenwood.scm

					
									+	 		"		#		'		*		2		4	 		4	  g  nameg  find-foundation CbRcCFHdeabG      h   `  ] 	$  C 	$  	64 5$  "  44 55$   4564 5$  "  	4	 	5$   4
44 5556 6     X      g  slot-id
	   g  filenamef  glenwood.scm

														 		' 			-!		0!		8!		;!		?		E#		M#			N$		X$			^%		j		p&		s&	/	v&	9	~&	/ &	 &		 (	 (		 	   g  nameg  to-foundations CcRCLHKFJIdf 
      h¨   Ò  ]
	$  C45$  "  r44 554455&  U44 554455$  "  '44 55$  4455"  "  $  	 6	 6   Ê      g  slot-id
	 ¥ g  tab-id	 ¥ g  t		S   g  filenamef  glenwood.scm
*
	+			+		-		-			.		!.		).		*/		-/		5/		9-			:0		=0	%	E0		F0		G1		J1	 	R1		S0		S0		a2		d2	%	l2		o2		s2		t3		w3	%	3	 3	 +	 4		 £5	1 ¥5	 #	 ¥	  g  nameg  check-a-tableau-with-single CfRCLM:KHFJId9g       hÐ   U  ]
	$  C45$  "   $  "  444 5554455&  a444 5554455$  "  -444 555$  4455	"  "  $  
 44 556 6M      g  slot-id
	 Ð g  tab-id	 Ð g  t		m ª  g  filenamef  glenwood.scm
7
	8			8		:		:			";		&:			,<		/<	!	2<	*	:<	!	;<		=<		>=		A=		I=		M:			N>		Q>	*	T>	3	\>	*	]>	%	_>		`>		a?		d?	 	l?		m>		m>		{@		~@	* @	3 @	* @	% @	 @	 @	 A	 A	% A	 ¢A	 ³8	 ¸B	 »B	$ ÃB	 ÇB		 ÎC	* ÐC	 1	 Ð	  g  nameg  check-a-tableau-pile CgRhCfg    hÀ     ] 	$  C 	$  	64 5$  "  @ $  "   	$  "   	
$  4 	5"  $   	64 5$  "  % 	$   	$  4 	5"  "  $   	6 6      g  slot-id
	 À g  t	0	Z g  t		B	W  g  filenamef  glenwood.scm
F
	G			G		I			G		J			K		'K			0L		0L		BM		BL		TN		^K			_O		pG		xP			yQ	 Q		 R	 Q		 S	 Q		 T	 ±G	 ¹U		 ¾W	 ÀW	 	 À  g  nameg  
to-tableau ChRC5ide     hÐ   ×  ]4	5$  "  54	5$  "  4		5$  "  4	
5$  y4	5$  "  84	5$  "   4	5$  "  4	5$  
45 C45$  C456C   Ï      g  slot-id
	 Í g  t			N g  t			K g  t		1	H g  t	Z ¢ g  t		o  g  t	    g  filenamef  glenwood.scm
Y
	Z	
		Z		[	
	Z		+\	
	1Z		?]	
	RZ		S^		Z^		Z^		h_		o_		o^		}`	 `	 ^	 a	 a	 ¦^	 ¨b	 ¬b	 ®b	 ±b	 ³c	 ¼^	 Ãd	 Ëd	 	 Í  g  nameg  empty-tableau? CjR`chj^        hP   ß   ]45   $   C45  $   C45  $   C4	5  $   C6       ×       g  t
		J g  t
		J g  t
	)	J g  t
	;	J  g  filenamef  glenwood.scm
h
	i		i		j		i		$k		)i		5l		;i		Jm	 		J
  g  nameg  get-hint CZR    h   W   ] C    O       g  filenamef  glenwood.scm
o
 		
  g  nameg  get-options CkR    h   o   ]C    g       g  options
		  g  filenamef  glenwood.scm
r
 		  g  nameg  apply-options ClR    h   S   ] C    K       g  filenamef  glenwood.scm
u
 		
  g  nameg  timeout CmR4nioi>  "  G  pi.iEiQiSiXi[iYiZikilimiNi6      h      g  filenamef  glenwood.scm		
		
  		 £	
 ¥		 ¨	
	
y	B
i	I
@	M

	Q
 	_
á	t
# ´
ø ¾
û Ã
 c Ç
"Ì Ð
%) ß
%ë ï
'Z ô
( ú
)Z
+Ú
-0
/K
1í*
5@7
7£F
:jY
;ºh
<*o
<²r
=u
=x
=jz
 %	=j
   C6 