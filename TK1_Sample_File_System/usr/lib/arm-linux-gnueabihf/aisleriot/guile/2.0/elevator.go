GOOF----LE-4-2.0«1      ] H 4      hH      ] g  guile¤	 ¤	g  process-use-modules¤	 ¤	 ¤	g  	aisleriot¤	g  	interface¤	 ¤		 ¤	
g  api¤	
 ¤	 ¤	g  initialize-playing-area¤	g  set-ace-low¤	g  make-standard-deck¤	g  shuffle-deck¤	g  add-normal-slot¤	g  DECK¤	g  add-blank-slot¤	g  add-carriage-return-slot¤	g  VERTPOS¤	e  0.5¤	g  HORIZPOS¤	g  
deal-cards¤											
												 ¤	g  deal-cards-face-up¤								 ¤	g  give-status-message¤	g  new-game¤	g  set-statusbar-message¤	g  get-stock-no-string¤	 g  string-append¤	!g  _¤	"f  Stock left:¤	#f   ¤	$g  number->string¤	%g  length¤	&g  	get-cards¤	'g  is-visible?¤	(g  button-pressed¤	)g  empty-slot?¤	*g  flip-top-card¤	+g  check-for-flips¤	,g  	get-value¤	-g  get-top-card¤	.g  king¤	/g  ace¤	0g  
droppable?¤	1g  add-to-score!¤	2g  move-n-cards!¤	3g  button-released¤	4 ¤	5g  	play-card¤	6g  	dealable?¤	7g  do-deal-next-cards¤	8g  click-to-move?¤	9g  button-clicked¤	:g  button-double-clicked¤	;g  	hint-move¤	<g  	playable?¤	=g  game-won¤	>g  get-hint¤	?g  game-continuable¤	@f  Deal a card¤	Ag  get-options¤	Bg  apply-options¤	Cg  timeout¤	Dg  set-features¤	Eg  droppable-feature¤	Fg  dealable-feature¤	Gg  
set-lambda¤C 5   hð*    ] 4	 >  "  G     h  4  ] 4>   "  G  4>   "  G  4>   "  G  4>   "  G  4>  "  G  4>  "  G  4>   "  G  4>  "  G  4>   "  G  	
 	
 4>   "  G  4>   "  G  4>  "  G  4>  "  G  4>   "  G  	
 	4>   "  G  4>   "  G  4>  "  G  4>  "  G  4>  "  G  4>   "  G  	
 	
 4>   "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>   "  G  	
 	4>   "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>   "  G  	
 	
 4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>   "  G  	
 	4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4
>  "  G  4
>  "  G  4>   "  G  		 C  ,      g  filenamef  elevator.scm
	
							#			3			C			U			X			]			f			v			y			~		 		 		 		 		 ¡	 	 ¢	 	 ¤	 	 ¥	!	 µ	"	 Å	#	 È	#	 Í	#	 Ö	$	 Ù	$	 Þ	$	 ç	&	 ú	'	 û	'	 ý	'	 þ	(		)		*	!	*	&	*	/	+	2	+	7	+	@	,	C	,	H	,	Q	.	d	/	e	/	g	/	k	0	l	0	n	0	o	1		2		2		2		3		3		3	¡	4	¤	4	©	4	²	5	µ	5	º	5	Ã	7	Ö	8	×	8	Ù	8	Ú	9	ê	:	í	:	ò	:	û	;	þ	;		;		<		<		<		=	 	=	%	=	.	>	1	>	6	>	?	@	R	A	S	A	U	A	Y	B	Z	B	\	B	]	C	`	C	e	C	n	D	q	D	v	D		E		E		E		F		F		F	¡	G	¤	G	©	G	²	H	µ	H	º	H	Ã	J	Ö	K	×	K	Ù	K	Ú	L	Ý	L	â	L	ë	M	î	M	ó	M	ü	N	ÿ	N		N		O		O		O		P	!	P	&	P	/	Q	2	Q	7	Q	@	R	C	R	H	R	Q	T	V	T	[	T	d	U	i	U	n	U	w	W		Y	 	
  g  nameg  new-game CR h   n   ] 45 6     f       g  filenamef  elevator.scm
	]
		^			^	 		
  g  nameg  give-status-message CR !"#$%&     h    ®   ] 45444
5556 ¦       g  filenamef  elevator.scm
	`
		a				a			a			a	"		b			b	!		b	)		b	!		b			a	 		
  g  nameg  get-stock-no-string CR' h   ¶   ]45$   CC  ®       g  slot-id
		 g  	card-list		  g  filenamef  elevator.scm
	d
		e			e		
	e			e			f			f	 			  g  nameg  button-pressed C(R)*       hè  8  ] 	$  4	5$  	6C 	$  44	5$  4	>  "  G  "   4	5$  	6C 	$  44	5$  4	>  "  G  "   4	5$  	6C 	$  44	5$  4	>  "  G  "   4	5$  	6C 	$  44	5$  4	>  "  G  "   4	5$  	6C 	$  44	5$  4	>  "  G  "   4	5$  	6C 	$  4	5$  	6C 	$  4	5$  	6C 	$  44	5$  4	>  "  G  "   4	5$  	6C 	$  44	5$  4	>  "  G  "   4	5$  	6C 	$  44	5$  4	>  "  G  "   4	5$  	6C 	$  44	5$  4	>  "  G  "   4	5$  	6C 	$  4	5$  	6C 	$  4	5$  	6C 	$  44	5$  4	>  "  G  "   4	5$  	
6C 	$  44	5$  4	
>  "  G  "   4	5$  		6C 	$  44	5$  4		>  "  G  "   4	5$  	6C 	$  4	5$  	6C 	$  4	
5$  	6C 	
$  44	5$  4	>  "  G  "   4		5$  	6C 		$  44	
5$  4	>  "  G  "   4	5$  	6C 	$  4		5$  	6C 	$  4	5$  	6C 	$  44	5$  4	>  "  G  "   4	5$  	6C 	$  4	5$  	6C 	$  4	5$  	6C 	$  4	5$  	6CC 0      g  slot-id
	ç  g  filenamef  elevator.scm
	h
		i				i			j			j				k		#	m			'	i		(	o		2	o		3	p		I	r		S	r		Y	s		`	u			d	i		e	w		o	w		p	x	 	z	 	z	 	{	 	}		 ¡	i	 ¢		 ¬		 ­ 	 Ã 	 Í 	 Ó 	 Ú 		 Þ	i	 ß 	 é 	 ê 	  	
 	 	 			i	 	& 	' 	= 	G 	M 	T 		X	i	Y 	c 		i 	p 		t	i	u 	 		 	 			i	 	 	  	² ¢	¼ ¢	Â £	É ¥		Í	i	Î §	Ø §	Ù ¨	ï ª	ù ª	ÿ «	 ­		
	i	 ¯	 ¯	 °	, ²	6 ²	< ³	C µ		G	i	H ·	R ·	S ¸	i º	s º	y »	 ½			i	 ¾	 ¾		 ¿	 Á		 	i	¡ Â	« Â		± Ã	¸ Å		¼	i	½ Ç	Ç Ç	È È	Þ Ê	è Ê	î Ë	õ Í		ù	i	ú Ï	 Ï	 Ð	 Ò	% Ò	+ Ó	2 Õ		6	i	7 ×	A ×	B Ø	X Ú	b Ú	h Û	o Ý		s	i	t Þ	~ Þ		 ß	 á			i	 â	 â		  ã	§ å		«	i	¬ ç	¶ ç	· è	Í ê	× ê	Ý ë	ä í		è	i	é ï	ó ï	ô ð	
 ò	 ò	 ó	! õ		%	i	& ö	0 ö		6 ÷	= ù		A	i	B ú	L ú		R û	Y ý		]	i	^ ÿ	h ÿ	i 							i		¥		«	²			¶	i	·
	Á
		Ç	Î		Ò	i	Ó	Ý		ã	 ¶	ç  g  nameg  check-for-flips C+R),-./       h¨   J  ]$  45$  C445545$  C445545$  C4455$  45"  $  C4455$  45CCC   B      g  
start-slot
	 ¥ g  	card-list	 ¥ g  end-slot		 ¥ g  t		+ £ g  t		K £ g  t		x £  g  filenamef  elevator.scm

			
										!		"		'		)		*		+		+		7		:		A		B		C		H		J		K		K		Y		\	#	c		d		h		k		p	"	r		s		x	 	 	" 	 	 	 	 	# 	  	 +	 ¥	  g  nameg  
droppable? C0R012+       h@   Î   ]4 5$  -4>  "  G  4 >  "  G   6CÆ       g  
start-slot
		@ g  	card-list		@ g  end-slot			@  g  filenamef  elevator.scm

					!		#"		>#	 		@	  g  nameg  button-released C3R)4'-,./1+     h@  Ó  ] 
$  4
5$  C
6 $  "  â4 5$  "  Ò44 55$  Á45$  "  ®445544 55$  "  445544 55$  "  \4455$  44 55"  $  "  &4455$  44 55"  "  $  +4	>  "  G  4
 >  "  G   6C   Ë      g  slot-id
	= g  t	r g  t	  g  t	 Ë  g  filenamef  elevator.scm
&
	(	
	
'		)		)			*	#	*		!,		%,			+-		5,			;.		>.		F.		J,			K/		T,			Z0		]0	 	d0		e1		h1	%	p1		q1		r0		r0	 2	 2	% 2	 2	 3	 3	  3	 2	 0	 ¨4	 «4	* ²4	 ³4	 ·4	 º5	 ½5	) Å5	 Æ5	 Ë0	 Û6	 Þ6	) å6	 æ6	 ê6	 í7	 ð7	* ø7	 ù7	'	9	":	(:	-:	;;	 =	=  g  nameg  	play-card C5R)    h   g   ] 4
5C      _       g  filenamef  elevator.scm
>
	?			?	 		

  g  nameg  	dealable? C6R5        h   g   ] 
6 _       g  filenamef  elevator.scm
A
	B	 		
  g  nameg  do-deal-next-cards C7R58      h(   ¦   ]"   645 $   $  C"ÿÿã"ÿÿß        g  slot-id
		'  g  filenamef  elevator.scm
F
	J		G		G		G		H		G	 		'  g  nameg  button-clicked C9R5 h      ] 6x       g  slot-id
		  g  filenamef  elevator.scm
L
	M	 		  g  nameg  button-double-clicked C:R)'-,./;< 	h    ] 	$  "  45$  C4 5$  "  Ã44 55$  ²445544 55$  "  445544 55$  "  \4455$  44 55"  $  "  &4455$  44 55"  "  $   6 6             g  
check-slot
		 g  t		 g  t	\ ñ g  t	  î g  t	 µ ë  g  filenamef  elevator.scm
O
	P	
	P		Q	
	!P		$S		.S	
	4T		7T		?T		CS	
	DU		GU	!	NU		OV		RV	&	ZV		[V		\U		\U		jW		mW	&	tW		uW		vX		yX	! X	 W	 U	 Y	  Y	+ Y	  Y	 ¡Y	 ¤Z	 §Z	* ¯Z	 °Z	 µU	 Å[	 È[	* Ï[	 Ð[	 Ô[	 ×\	  Ú\	+ â\	  ã\	 úS	]	
^		^	
 5		  g  nameg  	playable? C<R=>       h(      ] 4>   "  G  45 $  C6        x       g  filenamef  elevator.scm
`
	a		b		b		!c	 		!
  g  nameg  game-continuable C?R)       h   ]   ] 	6U       g  filenamef  elevator.scm
e
	f	 		
  g  nameg  game-won C=R)!@      h       ] 4
5$  C
45 C             g  filenamef  elevator.scm
h
	i		i		j		j		j		j	 		
  g  nameg  	dealable? C6R<6  h      ]4	5  $   C6 w       g  t
			  g  filenamef  elevator.scm
m
	n			n		o	 		
  g  nameg  get-hint C>R    h   W   ] C    O       g  filenamef  elevator.scm
q
 		
  g  nameg  get-options CAR    h   o   ]C    g       g  options
		  g  filenamef  elevator.scm
t
 		  g  nameg  apply-options CBR    h   S   ] C    K       g  filenamef  elevator.scm
w
 		
  g  nameg  timeout CCR4DiEiFi>  "  G  Gii(i3i9i:i?i=i>iAiBiCi0i6i6      ù       g  filenamef  elevator.scm		
		
	¡	]

	`
q	d
«	h
½
é
 &&
 ²>
!:A
"!F
"»L
&O
'C`
'Àe
(h
)2m
)¢q
**t
*w
*z
*ê|
 	*ê
   C6 