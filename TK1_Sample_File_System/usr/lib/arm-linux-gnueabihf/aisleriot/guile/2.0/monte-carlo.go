GOOF----LE-4-2.0%(      ] B 4    h      ] g  guile¤	 ¤	g  process-use-modules¤	 ¤	 ¤	g  	aisleriot¤	g  	interface¤	 ¤		 ¤	
g  api¤	
 ¤	 ¤	g  initialize-playing-area¤	g  set-ace-low¤	g  make-standard-deck¤	g  shuffle-deck¤	g  add-normal-slot¤	g  DECK¤	g  add-blank-slot¤	g  add-carriage-return-slot¤	g  deal-cards-face-up¤											
															 ¤	g  give-status-message¤	g  new-game¤	g  set-statusbar-message¤	g  get-stock-no-string¤	g  string-append¤	g  _¤	f  Stock left:¤	f   ¤	g  number->string¤	 g  length¤	!g  	get-cards¤	"g  empty-slot?¤	#g  button-pressed¤	$g  	adjacent?¤	%g  	get-value¤	&g  get-top-card¤	'g  
droppable?¤	(g  add-to-score!¤	)g  remove-card¤	*g  button-released¤	+g  deal-the-empties¤	,g  moving-over¤	-g  
deal-cards¤	.g  move-cards-up¤	/g  button-clicked¤	0g  button-double-clicked¤	1g  game-won¤	2g  get-hint¤	3g  game-continuable¤	4g  	get-score¤	5g  	hint-move¤	6g  horizontal-check¤	7g  vertical-check¤	8g  backslash-check¤	9g  slash-check¤	:f  Deal more cards¤	;g  empty?¤	<g  get-options¤	=g  apply-options¤	>g  timeout¤	?g  set-features¤	@g  droppable-feature¤	Ag  
set-lambda¤C 5  h !    ] 4	 >  "  G       hø  ½  ] 4>   "  G  4>   "  G  4>   "  G  4>   "  G  4>  "  G  4>   "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>   "  G  4>   "  G  4>   "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>   "  G  4>   "  G  4>   "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>   "  G  4>   "  G  4>   "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>   "  G  4>   "  G  4>   "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4	

>  "  G  4>   "  G  		 Cµ      g  filenamef  monte-carlo.scm
	
							#			3			C			U			e			h			m			v			y			~		 		 		 		 		 		  		 ©		 ¬		 ±		 º		 Ê	 	 Ú	!	 ê	"	 í	"	 ò	"	 û	#	 þ	#		#		$		$		$		%	 	%	%	%	.	&	1	&	6	&	?	'	O	(	_	)	o	*	r	*	w	*		+		+		+		,		,		,	¢	-	¥	-	ª	-	³	.	¶	.	»	.	Ä	/	Ô	0	ä	1	ô	2	÷	2	ü	2		3		3		3		4		4		4	'	5	*	5	/	5	8	6	;	6	@	6	I	7	Y	8	i	9	y	:	|	:		:		;		;		;		<		<	£	<	¬	=	¯	=	´	=	½	>	À	>	Å	>	Î	?	Ó	?	Ø	?	á	B	÷	D	 d	ø
  g  nameg  new-game CR        h   q   ] 45 6     i       g  filenamef  monte-carlo.scm
	F
		G			G	 		
  g  nameg  give-status-message CR !  h    ±   ] 45444
5556 ©       g  filenamef  monte-carlo.scm
	I
		J				J			J			J	"		K			K	!		K	)		K	!		K			J	 		
  g  nameg  get-stock-no-string CR"      h   ©   ] 
$  	4 5CC   ¡       g  slot-id
		 g  	card-list		  g  filenamef  monte-carlo.scm
	M
		N		
	N			O			O	 			  g  nameg  button-pressed C#R  hð     ]
 	$   	"  $  C 	$  C 	$   	
"  $  C $   	"  $  C $   	
"  $  C 	$   	"  $  C 	$  C 	$  	 	
CC {      g  slot1
	 ï g  slot2	 ï g  t		 ï g  t		. ï g  t		R ï g  t		t ï g  t	  ï g  t	 º ï g  t	 Í ï  	g  filenamef  monte-carlo.scm
	Q
			R		
	R			R			S			S			S			R		-	T		.	T		.	R		@	U		A	U		E	U		J	V		L	V		M	V		R	R		b	W		c	W		g	W		l	X		n	X		o	X		t	R	 	Y	 	Y	 	Y	 	Z	 	Z	 	Z	 	R	 ¨	[	 ©	[	 ­	[	 ²	\	 ´	\	 µ	\	 º	R	 Ì	]	 Í	]	 Í	R	 ß	^	 à	^	 ä	^	 é	_	 ë	_	 ì	_	 1	 ï	  g  nameg  	adjacent? C$R"$%&      hP     ] $  C
$  245$  C4 5$  454455CCC             g  
start-slot
		I g  	card-list		I g  end-slot			I  g  filenamef  monte-carlo.scm
	c
		d			d			e			d			f		 	d		#	g		/	d		0	h	
	5	h		7	h	
	8	i	
	;	i		C	i	
	D	h	 		I	  g  nameg  
droppable? C'R'()     h(   Ë   ]4 5$  4	5$  6CC  Ã       g  
start-slot
		& g  	card-list		& g  end-slot			&  g  filenamef  monte-carlo.scm
	k
		l			l			m			l		"	n	 		&	  g  nameg  button-released C*R"+        h@   á   ]
 	$  "  4
5$  C4
  5$   6C   Ù       g  slot
		= g  acc		= g  t			  g  filenamef  monte-carlo.scm
	p
		q	
		q			r	
	!	q		%	t		-	t	!	/	t		3	t		8	u		;	u	 		=	  g  nameg  deal-the-empties C+R+",-        h`   1  ] 	$  	64 5$   6
$  4   5$  
 6C 6 )      g  slot
		_ g  blanks		_ g  acc			_  g  filenamef  monte-carlo.scm
	w
		x				x			y			y				z			!	x		&	{		)	{	!	-	{			1	|			5	x		6	}		?	}	%	B	}		D	}		H	}			M	~		R	~		Y			_		 		_	  g  nameg  moving-over C,R,      h   e   ] 
6       ]       g  filenamef  monte-carlo.scm
 
		 	 			
  g  nameg  move-cards-up C.R.  h      ] 
$  6 C       g  slot-id
		  g  filenamef  monte-carlo.scm
 
	 		
 		 	 		  g  nameg  button-clicked C/R     h   z   ]C    r       g  slot-id
		  g  filenamef  monte-carlo.scm
 
 		  g  nameg  button-double-clicked C0R12 h(      ] 4>   "  G  45 $  C6        {       g  filenamef  monte-carlo.scm
 
	 		 		 		! 	 		!
  g  nameg  game-continuable C3R4    h   i   ] 	445 C     a       g  filenamef  monte-carlo.scm
 
	 		
 	 		
  g  nameg  game-won C1R"%&56      hp   V  ]
 	$  "  ;4 5$  "  *4 5$  "  44 5544 55$  
  6 	$   6C   N      g  slot-id
		m  g  filenamef  monte-carlo.scm
 
	 			 		 			 		 	 	 		 			$ 		. 			4 		7 		? 		@ 		C 		H 	*	J 		L 		M 		Q 		Y 		[ 			` 			d 		i 		k 		 		m  g  nameg  horizontal-check C6R"%&57 hh   9  ]4 5$  "  /4	 5$  "  44 5544	 55$   	 6 	$   6C     1      g  slot-id
		c  g  filenamef  monte-carlo.scm
 
	  		  			 ¡		 ¡	 	 ¡		   			& ¢		) ¢		1 ¢		2 £		5 £		< £	*	> £		@ £		A ¢		E  		O ¤		Q ¤			V ¥			Z  		_ ¦		a ¦		 		c  g  nameg  vertical-check C7R"%&58      hx   U  ]
 	$  "  ?4	 5$  "  ,4 5$  "  44 5544	 55$   	 6 	$   6C     M      g  slot-id
		s  g  filenamef  monte-carlo.scm
 ©
	 ª			 ª		 ª			 «		 «	 	 «		  ª			& ¬		0 ª			6 ­		9 ­		A ­		B ®		E ®		L ®	*	N ®		P ®		Q ­		U ª		_ ¯		a ¯			f °			j ª		o ±		q ±		 		s  g  nameg  backslash-check C8R"%&59  hx   Q  ] 	$  "  ?4	 5$  "  ,4 5$  "  44 5544	 55$   	 6 	$   6C     I      g  slot-id
		s  g  filenamef  monte-carlo.scm
 ´
	 µ			 µ		 µ			 ¶		 ¶	 	 ¶		  µ			& ·		0 µ			6 ¸		9 ¸		A ¸		B ¹		E ¹		L ¹	*	N ¹		P ¹		Q ¸		U µ		_ º		a º			f »			j µ		o ¼		q ¼		 		s  g  nameg  slash-check C9R":;        h`   2  ] 	$   $  4
5$  C
45 CC4 5$   6$  
45 C 6       *      g  slot
		Y g  yesblank		Y  g  filenamef  monte-carlo.scm
 ¿
	 À			 À		 Á			 Â		 Á		 Ã		# Ã		% Ã		( Ã		, Å			6 À		; Æ		> Æ			D À		F È		J È		L È		O È			U É		Y É	 		Y	  g  nameg  empty? C;R6789;     hP   â   ]45  $   C45  $   C45  $   C4	5  $   C6   Ú       g  t
		M g  t
		M g  t
	*	M g  t
	<	M  g  filenamef  monte-carlo.scm
 Ë
	 Ì		 Ì		 Í		 Ì		% Î		* Ì		6 Ï		< Ì		M Ð	 		M
  g  nameg  get-hint C2R h   Z   ] C    R       g  filenamef  monte-carlo.scm
 Ò
 		
  g  nameg  get-options C<R h   r   ]C    j       g  options
		  g  filenamef  monte-carlo.scm
 Õ
 		  g  nameg  apply-options C=R h   V   ] C    N       g  filenamef  monte-carlo.scm
 Ø
 		
  g  nameg  timeout C>R4?i@i>  "  G  Aii#i*i/i0i3i1i2i<i=i>i'i6         g  filenamef  monte-carlo.scm		
ø	
	F
	I
	\	M
Þ	Q
a	c
n	k
¬	p
\	w
è 
 
- 
î 
| 
a 
 
 ©
ì ´
 ¿
í Ë
 ] Ò
 å Õ
!Q Ø
!R Û
! Ý
 	!
   C6 