GOOF----LE-4-2.0B      ] 3 4   h      ] g  guile¤	 ¤	g  process-use-modules¤	 ¤	 ¤	g  	aisleriot¤	g  	interface¤	 ¤		 ¤	
g  api¤	
 ¤	 ¤	g  initialize-playing-area¤	g  set-ace-low¤	g  make-standard-deck¤	g  shuffle-deck¤	g  add-extended-slot¤	g  down¤	g  add-carriage-return-slot¤	g  VERTPOS¤	e  0.75¤	g  HORIZPOS¤	g  add-normal-slot¤	g  DECK¤	g  deal-cards-face-up¤	
										
	 ¤	
		 ¤	g  new-game¤	g  empty-slot?¤	g  length¤	g  button-pressed¤	 g  	get-value¤	!g  get-top-card¤	"g  
droppable?¤	#g  add-to-score!¤	$g  remove-card¤	%g  button-released¤	&g  button-clicked¤	'g  button-double-clicked¤	(g  game-won¤	)g  get-hint¤	*g  game-continuable¤	+g  
check-slot¤	,g  	hint-move¤	-g  get-options¤	.g  apply-options¤	/g  timeout¤	0g  set-features¤	1g  droppable-feature¤	2g  
set-lambda¤C 5 h  ¸   ] 4	 >  "  G       hÈ  |  ] 4>   "  G  4>   "  G  4>   "  G  4>   "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>   "  G  	 4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  
¸ 
4>  "  G  4	>  "  G  4	>  "  G  4	>  "  G  4	>  "  G  4	>  "  G  		 C     t      g  filenamef  fourteen.scm
	
							#			3			C			F			M			V			Y			`			i			l			s			|					 		 		 		 		 ¢		 ¥		 ¬		 µ		 È	 	 É	 	 Ë	 	 Ì	"	 Ï	"	 Ö	"	 ß	#	 â	#	 é	#	 ò	$	 õ	$	 ü	$		%		%		%		&		&	"	&	+	'	.	'	5	'	C	)	E	)	F	*	X	,	^	,	c	,	l	-	r	-	w	-		.		.		.		/		/		/	¨	0	®	0	³	0	Â	2	 A	Ã
  g  nameg  new-game CR h(   ®   ]4 5$  "  	45$  CC     ¦       g  slot-id
		# g  	card-list		#  g  filenamef  fourteen.scm
	5
		6			6			7			7			6	 		#	  g  nameg  button-pressed CR !     h8   û   ]45$  C $  C	445545C     ó       g  
start-slot
		3 g  	card-list		3 g  end-slot			3  g  filenamef  fourteen.scm
	;
		<			<			=			<			>		 	>		(	>		)	?		.	?		0	?		1	>		2	>	 		3	  g  nameg  
droppable? C"R"#$        h(   È   ]4 5$  4	5$  6CC  À       g  
start-slot
		& g  	card-list		& g  end-slot			&  g  filenamef  fourteen.scm
	A
		B			B			C			B		"	D	 		&	  g  nameg  button-released C%R   h   o   ]C    g       g  slot-id
		  g  filenamef  fourteen.scm
	F
 		  g  nameg  button-clicked C&R    h   v   ]C    n       g  slot-id
		  g  filenamef  fourteen.scm
	I
 		  g  nameg  button-double-clicked C'R()       h   s   ] 45 $  C6        k       g  filenamef  fourteen.scm
	L
		M			M			N	 		
  g  nameg  game-continuable C*R    h     ] 4
5$  45$  {4	5$  n4	5$  a4	5$  T4	5$  G4	5$  :4	5$  -4	5$   4		5$  4	
5$  	6CCCCCCCCCCC         g  filenamef  fourteen.scm
	P
		Q			Q			R			Q			S		!	Q		"	T		,	Q		-	U		7	Q		8	V		B	Q		C	W		M	Q		N	X		X	Q		Y	Y		c	Q		d	Z		n	Q		o	[		y	Q			\	 	 
  g  nameg  game-won C(R+ !,   h   y  ]
4 5$   	
$   	 6C45$  "  	44 554455$  )	$  	 6 	
$   	 6C 6  q      g  slot1
	  g  slot2	  g  t		+	X  g  filenamef  fourteen.scm
	^
		_			_			`	
		`			a		 	a	"	"	a	
	%	c		+	c	
	;	d		>	d	'	F	d		G	e		J	e	'	R	e		S	d		T	d		U	d		\	c		a	f		e	f	
	l	g	 	n	g		s	h		w	h		|	i	 	i	* 	i	 	k	
 	 	  g  nameg  
check-slot C+R+      h   [   ] 
6S       g  filenamef  fourteen.scm
	m
		n	 		
  g  nameg  get-hint C)Rh   V   ] C    N       g  filenamef  fourteen.scm
	p
 		
  g  nameg  get-options C-R     h   n   ]C    f       g  options
		  g  filenamef  fourteen.scm
	s
 		  g  nameg  apply-options C.R     h   R   ] C    J       g  filenamef  fourteen.scm
	v
 		
  g  nameg  timeout C/R40i1i>  "  G  2iii%i&i'i*i(i)i-i.i/i"i6       °       g  filenamef  fourteen.scm		
	
y	5
Æ	;
Ó	A
Z	F
é	I
		L
G	P
l	^
æ	m
Q	p
Ù	s
E	v
F	y
	{
 	
   C6 