GOOF----LE-4-2.0Ù"      ] > 4        h       ] g  guile¤	 ¤	g  process-use-modules¤	 ¤	 ¤	g  	aisleriot¤	g  	interface¤	 ¤		 ¤	
g  api¤	
 ¤	 ¤	g  initialize-playing-area¤	g  set-ace-low¤	g  make-standard-deck¤	g  shuffle-deck¤	g  add-partially-extended-slot¤	g  right¤	g  HORIZPOS¤	g  add-carriage-return-slot¤	g  deal-cards-face-up-from-deck¤	g  DECK¤	



																																					
	
	
	
								 4¤	g  freeze-slots-if-complete¤	
										
		 ¤	g  new-game¤	g  length¤	g  	get-value¤	g  check-same-value-list¤	g  flip-top-card¤	g  add-to-score!¤	 g  freeze-slot¤	!g  is-visible?¤	"g  button-pressed¤	#g  	get-cards¤	$g  freeze-if-complete¤	%g  move-n-cards!¤	&g  complete-transaction¤	'g  empty-slot?¤	(g  
droppable?¤	)g  button-released¤	*g  button-clicked¤	+g  button-double-clicked¤	,g  game-won¤	-g  get-hint¤	.g  	game-over¤	/g  done-or-empty¤	0g  get-top-card¤	1g  check-number¤	2g  check-a-slot¤	3g  _¤	4f  	something¤	5f  an empty slot¤	6g  	hint-move¤	7g  check-slots¤	8g  get-options¤	9g  apply-options¤	:g  timeout¤	;g  set-features¤	<g  droppable-feature¤	=g  
set-lambda¤C 5   hp  ÿ   ] 4	 >  "  G   h(  $  ] 4>   "  G  4>   "  G  4>   "  G  4>   "  G  4	>  "  G   4	>  "  G   4	>  "  G   4	>  "  G  4>   "  G  4	>  "  G   4	>  "  G   4	>  "  G   4	>  "  G  4>   "  G  4	>  "  G   4	>  "  G   4	>  "  G   4	>  "  G  4>   "  G  4	>  "  G   4	>  "  G   4	>  "  G  4>   "  G  4	
>  "  G  4>  "  G  		 C            g  filenamef  
pileon.scm
	
							#			3			C			F			O			Z			\			]			`			i			t			v			w			z		 		 		 		 		 		 		 ¦		 ¶	!	 ¹	!	 Â	!	 Í	"	 Ï	"	 Ð	#	 Ó	#	 Ü	#	 ç	$	 é	$	 ê	%	 í	%	 ö	%		&		&		'		'		'		(	)	*	,	*	5	*	@	+	B	+	C	,	F	,	O	,	Z	-	\	-	]	.	`	.	i	.	t	/	v	/	w	0	z	0		0		1		3		3	¨	3	³	4	µ	4	¶	5	¹	5	Â	5	Í	6	Ï	6	Ð	7	Ó	7	Ü	7	å	8	õ	:	û	;	 	:			=		=		=	!	?	 T	"
  g  nameg  new-game CR       h8   Þ   ]4 5	$  C4 54 5$   6C       Ö       g  	card-list
		1  g  filenamef  
pileon.scm
	A
		B				B			B			D			D			D			D	)	 	D	4	#	D	)	$	D	
	(	D		-	E	!	/	E	
 		1  g  nameg  check-same-value-list CR       h    z   ]4 >  "  G  	6      r       g  slot-id
		  g  filenamef  
pileon.scm
	H
		I			J	 		  g  nameg  freeze-slot C R!   h   ¤   ]45$  6C         g  slot-id
		 g  	card-list		  g  filenamef  
pileon.scm
	L
		M			M			N			N	 			  g  nameg  button-pressed C"R#      hH   Á   ]44 55	$  *44 55$  4 >  "  G  "   "   C       ¹       g  slot-id
		A  g  filenamef  
pileon.scm
	P
		Q	
		Q			Q	
		Q			Q			R			R		!	R		%	Q		&	S	 		A  g  nameg  freeze-if-complete C$R$    h    ­   ] (  C4 5$   6C ¥       g  slots
		  g  filenamef  
pileon.scm
	V
		W			X			X			X			W			Y	!		Y	 			  g  nameg  freeze-slots-if-complete CR%$        h    ³   ]4 >  "  G  6  «       g  
start-slot
		 g  	card-list		 g  end-slot			  g  filenamef  
pileon.scm
	[
		\			]	 			  g  nameg  complete-transaction C&R'#      hX   @  ] $  C45$  "  445545$  445545	CC8      g  
start-slot
		X g  	card-list		X g  end-slot			X g  t			:  g  filenamef  
pileon.scm
	_
		`			`			a			a		"	b		%	b	 	,	b		.	b		/	c		4	c		6	c		7	b		>	`		?	d		B	d		J	d		K	d	+	R	d	
	U	d	 		X	  g  nameg  
droppable? C(R(&     h    ¶   ]4 5$  
 6C   ®       g  
start-slot
		 g  	card-list		 g  end-slot			  g  filenamef  
pileon.scm
	f
		g			g			h	 			  g  nameg  button-released C)R     h   m   ]C    e       g  slot-id
		  g  filenamef  
pileon.scm
	j
 		  g  nameg  button-clicked C*R      h   q   ]C    i       g  slot
		  g  filenamef  
pileon.scm
	l
 		  g  nameg  button-double-clicked C+R,-    h   j   ] 45 $  C6        b       g  filenamef  
pileon.scm
	n
		o			o			p	 		
  g  nameg  	game-over C.R'!# h(   ²   ]	4 5$  C44 55C     ª       g  slot-id
		# g  t			#  g  filenamef  
pileon.scm
	r
		s				s			t			t			t		!	t		"	t	 			#  g  nameg  done-or-empty C/R/     hÀ   F  ] 4
5$  ®45$  ¢4	5$  4	5$  4	5$  {4	5$  n4	5$  a4	5$  T4	5$  G4		5$  :4	
5$  -4	5$   4	5$  4	5$  	6CCCCCCCCCCCCCC    >      g  filenamef  
pileon.scm
	v
		w			w			x			w			y		!	w		"	z		,	w		-	{		7	w		8	|		B	w		C	}		M	w		N	~		X	w		Y			c	w		d 		n	w		o 		y	w		z 	 	w	  	 	w	  	 	w	   	 	 ¼
  g  nameg  game-won C,R#0   hÀ   	  ]"  ;44 55	$  &44 5544 55$  C	CC"  >44 55	$  '44 5544 55$  "ÿÿ	C"ÿÿ44 55$  %44 5544 55$  "ÿÿC"ÿÿ        g  slot-id
	 ¾  g  filenamef  
pileon.scm
 
	 		
 		 		 		 		 		 	!	% 		& 		) 	)	0 	!	5 		6 		: 			A 		F 		I 		Q 		T 		X 		Y 		\ 	!	d 		e 		h 	(	o 	!	s 		t 		x 		  	  	  	  	  	  	  	  	! ¡ 	 ¢ 	 ¥ 	' ¬ 	! ¯ 	 ° 	 ´ 		 .	 ¾  g  nameg  check-number C1R2'345#06    hà   §  ]	$  C $   64 5$  	4545 C45$  	4545 C	4455$   644	5544	 55$  :	4455$   	44556
 6 6       g  slot-id
	 ß g  number-to-move	 ß g  to-slot		 ß  g  filenamef  
pileon.scm
 
	 		 		 	
	 		 	/	! 	
	" 		, 	
	/ 		3 		5 		6 	(	: 	*	< 	(	? 		A 		K 	
	N 		R 		T 		U 	(	Y 	*	[ 	(	^ 		b  		e  		m  		n  		r 	
	{ ¡	6	} ¡		~ ¢	  ¢	  ¢	  £	  £	  £	  ¢	  	
  ¤	, ¢ ¤	4 ª ¤	, « ¤	' ¬ ¤	 ° ¤	 ¹ §	+ ¼ §	3 Ä §	+ Å §	& Æ ¦	# É ¨	# Ë ¥	 Ô ©	 Ý ª	9 ß ª	 9	 ß	  g  nameg  check-a-slot C2R217    h0   ×   ]
 	$  C4 4 55$  C 
6Ï       g  slot-id
		0 g  to-slot		0 g  t			0  g  filenamef  
pileon.scm
 ¬
	 ­		 ­		 ¯	
	 ¯	 	 ¯	
	 ¯		- °		0 °	
 
		0	  g  nameg  check-slots C7R7        h   [   ] 
6S       g  filenamef  
pileon.scm
 ²
	 ³	 		
  g  nameg  get-hint C-Rh   U   ] C    M       g  filenamef  
pileon.scm
 µ
 		
  g  nameg  get-options C8R      h   m   ]C    e       g  options
		  g  filenamef  
pileon.scm
 ·
 		  g  nameg  apply-options C9R      h   Q   ] C    I       g  filenamef  
pileon.scm
 ¹
 		
  g  nameg  timeout C:R4;i<i>  "  G  =ii"i)i*i+i.i,i-i8i9i:i(i6÷       g  filenamef  
pileon.scm		
	
Á	A
u	H
G	L
	l	P

P	V
>	[
ó	_
á	f
h	j
ô	l
	n
}	r
	v
| 
* 
J ¬
Æ ²
0 µ
¸ ·
$ ¹
% »
p ½
 	p
   C6 