GOOF----LE-4-2.0       ] D 4  h      ] g  guile€	 €	g  process-use-modules€	 €	 €	g  	aisleriot€	g  	interface€	 €		 €	
g  api€	
 €	 €	g  initialize-playing-area€	g  set-ace-low€	g  make-joker-deck€	g  shuffle-deck€	g  add-extended-slot€	g  down€	g  add-carriage-return-slot€	g  add-normal-slot€	g  DECK€	g  deal-cards-face-up€	
					
					
					
					
						 $€	g  give-status-message€	g  new-game€	g  set-statusbar-message€	g  get-stock-no-string€	g  string-append€	g  _€	f  Stock left:€	f   €	 g  number->string€	!g  length€	"g  	get-cards€	#g  	get-value€	$g  joker€	%g  values-match?€	&g  ace€	'g  jack€	(g  queen€	)g  king€	*g  	score-for€	+g  empty-slot?€	,g  get-top-card€	-g  can-move-from€	.g  move-possible€	/g  button-pressed€	0g  add-to-score!€	1g  move-n-cards!€	2g  button-released€	3g  
droppable?€	4	 €	5g  
deal-cards€	6g  button-clicked€	7g  button-double-clicked€	8g  game-won€	9g  	game-over€	:g  	hint-move€	;g  hint-move-from€	<f  Deal a card from the deck€	=g  get-hint€	>g  get-options€	?g  apply-options€	@g  timeout€	Ag  set-features€	Bg  droppable-feature€	Cg  
set-lambda€C 5hΘ  ξ   ] 4	 >  "  G     h8    ] 4>   "  G  4>   "  G  4>   "  G  4>   "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>   "  G  4>   "  G  4	>  "  G  4>  "  G  4
	>  "  G  4>   "  G  		 C        g  filenamef  thieves.scm
	
							#			3			C			F			M			V			Y			`			i			l			s			|					 		 		 		 		 ’		 ₯		 ¬		 ΅		 Έ		 Ώ		 Θ	 	 Ψ	!	 θ	"	 ϊ	#	 ύ	#		#		$		$		$		'	5	)	 &	6
  g  nameg  new-game CR     h   m   ] 45 6     e       g  filenamef  thieves.scm
	+
		,			,	 		
  g  nameg  give-status-message CR !"      h    ­   ] 45444	5556₯       g  filenamef  thieves.scm
	.
		/				/			/			/	"		0			0	!		0	)		0	!		0		 	/	 		 
  g  nameg  get-stock-no-string CR#$        hX     ]
4 5$  C45$  C4 545$  C4 545C      g  c1
		X g  c2		X g  t			X g  t		!	X g  t		<	X  g  filenamef  thieves.scm
	2
		3			3			3			4		!	4		!	3		-	5		4	5		5	5	 	<	5		<	3		H	6		O	6		V	6		W	6	 		X	  g  nameg  values-match? C%R&'()    h   p   ] &  	C 	&  	C 	&  	C 	&  	C 	&  	C 	&  	C 	&  	C 	&  	C 		&  	C 	
&  	C &  	C &  	C &  	C
C     h       g  card
	   g  filenamef  thieves.scm
	8
	
	9	 	   g  nameg  	score-for C*R+%,   h0   £   ]4 5$  C4	5$  C4 54	56         g  where
		.  g  filenamef  thieves.scm
	H
		I			I			J			I			K		&	K	+	.	K	 			.  g  nameg  can-move-from C-R-    hx     ]4
5  $   C45  $   C4	5  $   C4	5  $   C4	5  $   C4	5  $   C	6            g  t
		r g  t
		r g  t
	+	r g  t
	=	r g  t
	O	r g  t
	a	r  g  filenamef  thieves.scm
	M
		N			N			O			N		%	P		+	N		7	Q		=	N		I	R		O	N		[	S		a	N		r	T	 		r
  g  nameg  move-possible C.R+!        h(   ΅   ] 	$  4 5$  C45CC    ­       g  slot-id
		$ g  	card-list		$  g  filenamef  thieves.scm
	V
		W			W			X			W			Y	
	!	Y	 		$	  g  nameg  button-pressed C/R%,0*#1        hH     ]	$  744	55$  "444555$  
 	6CCC          g  
start-slot
		D g  	card-list		D g  end-slot			D  g  filenamef  thieves.scm
	[
		\			\			]			]			]	&		]			\			^		"	^		%	^	!	*	^	,	,	^	!	.	^		0	^		4	\		>	_	 		D	  g  nameg  button-released C2R%,      h    Β   ]	$  4	56C     Ί       g  
start-slot
		 g  	card-list		 g  end-slot			  g  filenamef  thieves.scm
	a
		b			b			c			c	&		c	 			  g  nameg  
droppable? C3R+4%,0*#5 
    hx   .  ] 	&  4 5$  C	6 	$  J4 5$  C44 54	55$  $4444 5555$  	 6CCC    &      g  slot-id
		t  g  filenamef  thieves.scm
	f
	
	g			h			h			i	!		i		$	j		(	j		)	k		3	j		6	l		9	l		@	l	1	H	l		L	j		M	m		P	m		S	m	%	V	m	0	^	m	%	`	m		b	m		f	j		l	n		n	n	 		t  g  nameg  button-clicked C6R6 h   z   ] 6r       g  slot
		  g  filenamef  thieves.scm
	p
		q	 		  g  nameg  button-double-clicked C7R+     hX   Ί   ] 4
5$  F45$  :4	5$  -4	5$   4	5$  4	5$  	6CCCCCC    ²       g  filenamef  thieves.scm
	s
		t			t			u			t			v		!	t		"	w		,	t		-	x		7	t		8	y		B	t		H	z	 		T
  g  nameg  game-won C8R8+.       h0      ] 4>   "  G  45 $  C4	5$  6 C  ~       g  filenamef  thieves.scm
	|
		}			~			~		 		( 		, 	 		.
  g  nameg  	game-over C9R+%,:   hX   ν   ]4 5$  "  +4 5$  "  44 54	55$  C 	6     ε       g  where
		S g  t			D g  t			A  g  filenamef  thieves.scm
 
	 	
		 		 	
	 		+ 		. 		5 	3	= 		> 	
	H 		S 	 		S  g  nameg  hint-move-from C;R;<      h   e  ]4
5  $   C45  $   C4	5  $   C4	5  $   C4	5  $   C4	5  $   C4	5  $   C
45 C      ]      g  t
	  g  t
	  g  t
	+  g  t
	=  g  t
	O  g  t
	a  g  t
	s   g  filenamef  thieves.scm
 
	 		 		 		 		% 		+ 		7 		= 		I 		O 		[ 		a 		m 		s 	  	  	  	  	 	 
  g  nameg  get-hint C=R      h   V   ] C    N       g  filenamef  thieves.scm
 
 		
  g  nameg  get-options C>R     h   n   ]C    f       g  options
		  g  filenamef  thieves.scm
 
 		  g  nameg  apply-options C?R     h   R   ] C    J       g  filenamef  thieves.scm
 
 		
  g  nameg  timeout C@R4AiBi>  "  G  Cii/i2i6i7i9i8i=i>i?i@i3i6       ζ       g  filenamef  thieves.scm		
	
	+
	.
	2
;	8
&	H
	Θ	M

ΐ	V
J	[
E	a
	f
₯	p
Ν	s
‘	|
  
 
 
	 
u 
v 
Α 
 	Α
   C6 