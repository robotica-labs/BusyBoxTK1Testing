GOOF----LE-4-2.0S'      ] N 4        h@      ] g  guile¤	 ¤	g  process-use-modules¤	 ¤	 ¤	g  	aisleriot¤	g  	interface¤	 ¤		 ¤	
g  api¤	
 ¤	 ¤	g  	get-value¤	g  is-pair?¤	g  
find-match¤	g  	pair-deck¤	g  list-ref¤	g  	list-set!¤	g  find-hole-helper¤	g  	find-hole¤	g  find-deepest¤	g  map¤	g  max¤	g  +¤	g  aisleriot-random¤	g  select-first¤	g  select-second¤	g  empty-slot?¤	g  list->vector¤	g  	add-card!¤	g  make-visible¤	 g  remove-card¤	!g  deal-cards-in-pairs¤	"g  vector->list¤	#g  initialize-playing-area¤	$g  set-ace-low¤	%g  make-standard-deck¤	&g  shuffle-deck¤	'g  DECK¤	(g  add-extended-slot¤	)g  right¤	*g  add-carriage-return-slot¤	+g  down¤	,g  add-blank-slot¤	-g  add-normal-slot¤	.	
 ¤	/	 ¤	0		 ¤	1		 ¤	2		 ¤	3		 ¤	4		 ¤	5		 ¤	6		 ¤	7			 ¤	8		
 ¤	9./012345678 ¤	:g  new-game¤	;g  length¤	<g  button-pressed¤	=g  get-top-card¤	>g  
droppable?¤	?g  add-to-score!¤	@g  button-released¤	Ag  button-clicked¤	Bg  button-double-clicked¤	Cg  game-won¤	Dg  get-hint¤	Eg  game-continuable¤	Fg  check-for-pairs¤	Gg  	hint-move¤	Hg  get-options¤	Ig  apply-options¤	Jg  timeout¤	Kg  set-features¤	Lg  droppable-feature¤	Mg  
set-lambda¤C 5   h     ] 4	 >  "  G   hP   8  ]"4 545	$  C	$  		"  $  C	C 0      g  card1
		O g  card2		O g  val1				O g  val2			O g  sum			O g  t		 	O g  t		>	O  g  filenamef  gay-gordons.scm
	
																				 			 			0			4			9			>			N	 	 		O	  g  nameg  is-pair? CR     h(   Î   ] (  C4 5$  C 6Æ       g  target
		( g  list		( g  n			(  g  filenamef  gay-gordons.scm
	'
		(			)			)			)			(		#	*		&	*	)	(	*	 
		(	  g  nameg  
find-match CR   h`   d  ]! (  C (   6  4
5454>  "  G  45C \      g  cards
		_ g  first		_ g  	remainder		 	_ g  n		,	_ g  nth		7	_  g  filenamef  gay-gordons.scm
	/
		0				0			1			0			1	'		1			2			2		 	3		 	2		#	4		,	2		/	5		7	2		:	6		C	6	$	H	6		U	7	#	Z	7	.	\	7	#	]	7		^	7	 		_  g  nameg  	pair-deck CR   hX   2  ] (  	ÿC 
$  "  $   6
$  C 6      *      g  slots
		R g  n		R g  i			R g  avoid			R g  t			&  g  filenamef  gay-gordons.scm
	A
		B			C			C			C			C		#	D		*	B		/	E		2	E	&	8	E		<	F		@	B		H	G		K	G	)	N	G	1	R	G	 		R	  g  nameg  find-hole-helper CR     h      ] 
6          g  slots
		 g  n		 g  avoid			  g  filenamef  gay-gordons.scm
	I
		J	 			  g  nameg  	find-hole CR  h    »   ] $  
C4 5C      ³       g  holes
		 g  deepest		  g  filenamef  gay-gordons.scm
	M
		N				N			N			P			P			P			P	 				  g  nameg  find-deepest CR     h   Z   ] C  R       g  x
		  g  filenamef  gay-gordons.scm
	W			W	! 		   Ch   j   ] 
$  
CC b       g  x
		  g  filenamef  gay-gordons.scm
	Z			Z	,		Z	)		Z	% 		   C 	    hX   N  ])4 54?4?44 5?45	$  6 	ÿ6  F      g  slots
		V g  holes		V g  deepest			V g  total			V g  n		-	V g  r		6	V  g  filenamef  gay-gordons.scm
	V
		W			W			X			W			Y			W		 	Z		#	Z		-	Z		-	W		0	[		6	W		?	\		@	\		D	\		L	^		V	_	 		V  g  nameg  select-first CRh   j   ] 
$  
CC b       g  x
		  g  filenamef  gay-gordons.scm
	b			b	-		b	*		b	& 		   C       h(   è   ]44 5?45 6  à       g  slots
		& g  first		& g  n			& g  r			&  g  filenamef  gay-gordons.scm
	a
		b			b			b			b			c			c			c			b		&	d	 		&	  g  nameg  select-second CR !" 
      h°   G  ]:4 5$  C454545££444 55>  "  G  444 55>  "  G   ¤ ¤ 4	56     ?      g  source
	 « g  slots	 « g  vslots		 « g  a		 « g  b		* « g  a-pair		1 « g  b-pair		8 « g  a-slot		> « g  b-slot		D «  	g  filenamef  gay-gordons.scm
	h
		i			i			j			o			o			p			o		"	q		*	o		1	r		1	o		8	s		8	o		=	t		>	o		C	u		D	o		G	w		L	w		O	w	'	W	w		\	w		e	x		j	x		m	x	'	u	x		z	x	 	z	& 	z	# 	z	 	z	 	{	& 	{	# 	{	 	{	 £	|	# «	|	 '	 «	  g  nameg  deal-cards-in-pairs C!R#$%&'()*+,-!9      h   o  ] 4>   "  G  4>   "  G  4>   "  G  4>   "  G  45 4>  "  G  4	>   "  G  4
>  "  G  4
>  "  G  4
>  "  G  4
>  "  G  4
>  "  G  4
>  "  G  4
>  "  G  4
>  "  G  4
>  "  G  4
>  "  G  4>   "  G  4>   "  G  4>   "  G  4>   "  G  4>   "  G  4>   "  G  4>   "  G  4>   "  G  4>   "  G  4>   "  G  4>  "  G  4	>  "  G  	
	 C      g      g  filenamef  gay-gordons.scm
	~
				 		# 		3 		C 		K 		L 		O 		V 		_ 		o 		r 		y 	  	  	  	  	  	  	 ¨ 	 « 	 ² 	 » 	 ¾ 	 Å 	 Î 	 Ñ 	 Ø 	 á 	 ä 	 ë 	 ô 	 ÷ 	 þ 	 	
 	 	 	 	$ 	- 	= 	M 	] 	m 	} 	 	 	­ 	½ 	Í 	ß  	å  	ê  	ù ©	 9	ú
  g  nameg  new-game C:R;      h    ®   ]4 5$  C45C       ¦       g  slot-id
		 g  	card-list		  g  filenamef  gay-gordons.scm
 «
	 ¬		 ¬		 ­	
	 ­	 			  g  nameg  button-pressed C<R=     h¨   @  ]45$  C $  C	445545$  C	4455$  	45"  $  C4455	$  (45	$  445545CCC8      g  
start-slot
	 ¨ g  	card-list	 ¨ g  end-slot		 ¨ g  t		2 ¨ g  t		` ¨  g  filenamef  gay-gordons.scm
 ¯
	 °		 °		 ±		 °		 ²		  ²		( ²		) ³		. ³		0 ³		1 ²		2 ²		2 ²		@ ´		C ´	!	K ´		L ´		P ´		S µ		X µ	!	Z µ		[ µ		` ²		l ¶		o ¶		w ¶		z ¶		~ ¶		 ·	  ·	  ·	  ·	  ¶	  ¸	  ¸	#  ¸	  ¹	  ¹	# ¡ ¹	 ¢ ¸	 £ ¸	 +	 ¨	  g  nameg  
droppable? C>R> ?   h(   Ñ   ]4 5$  45$  	6CC  É       g  
start-slot
		& g  	card-list		& g  end-slot			&  g  filenamef  gay-gordons.scm
 »
	 ¼		 ¼		 ½		 ¼		" ¾	 		&	  g  nameg  button-released C@R  h   s   ]C    k       g  slot-id
		  g  filenamef  gay-gordons.scm
 À
 		  g  nameg  button-clicked CARh   z   ]C    r       g  slot-id
		  g  filenamef  gay-gordons.scm
 Ã
 		  g  nameg  button-double-clicked CBRCD   h   z   ] 45 $  C6        r       g  filenamef  gay-gordons.scm
 Æ
	 Ç		 Ç		 È	 		
  g  nameg  game-continuable CER     h     ] 4
5$  z45$  n4	5$  a4	5$  T4	5$  G4	5$  :4	5$  -4	5$   4	5$  4		5$  	
6CCCCCCCCCC      g  filenamef  gay-gordons.scm
 Ê
	 Ë		 Ë		 Ì		 Ë		 Í		! Ë		" Î		, Ë		- Ï		7 Ë		8 Ð		B Ë		C Ñ		M Ë		N Ò		X Ë		Y Ó		c Ë		d Ô		n Ë		t Õ	 	 
  g  nameg  game-won CCRF=G  hx   ?  ]
 	
$  C4 5$  "  	$   	 645$  "  44 5455$  	 6 6      7      g  slot1
		r g  slot2		r g  t			)  g  filenamef  gay-gordons.scm
 ×
	 Ø			 Ø		 Ú		 Ú			& Û		- Ø		2 Ü		7 Ü	&	9 Ü			: Ý		D Ý			J Þ		M Þ		T Þ	-	\ Þ		` Ø		i ß			p à	#	r à	 		r	  g  nameg  check-for-pairs CFRF        h   `   ] 
6X       g  filenamef  gay-gordons.scm
 â
	 ã	 		
  g  nameg  get-hint CDR   h   Z   ] C    R       g  filenamef  gay-gordons.scm
 å
 		
  g  nameg  get-options CHR h   r   ]C    j       g  options
		  g  filenamef  gay-gordons.scm
 è
 		  g  nameg  apply-options CIR h   V   ] C    N       g  filenamef  gay-gordons.scm
 ë
 		
  g  nameg  timeout CJR4KiLi>  "  G  Mi:i<i@iAiBiEiCiDiHiIiJi>i6   	      g  filenamef  gay-gordons.scm		
³	
Á	'
	/
=	A
ø	I
æ	M

¡	V
S	a
r	h
	~
ù «
û ¯
 »
 À
% Ã
Í Æ
 Ê
R ×
Ó â
E å
Í è
 9 ë
 : î
  ð
 	 
   C6 