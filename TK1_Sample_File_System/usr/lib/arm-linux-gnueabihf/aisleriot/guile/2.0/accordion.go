GOOF----LE-4-2.0!      ] E 4 h      ] g  guile¤	 ¤	g  process-use-modules¤	 ¤	 ¤	g  	aisleriot¤	g  	interface¤	 ¤		 ¤	
g  api¤	
 ¤	 ¤	
							 	¤	g  row1¤				
							 	¤	g  row2¤										 	¤	g  row3¤							 	!	"	# 	¤	g  row4¤		$	%	&	'	(	)	*	+	, 	¤	g  row5¤		-	.	/	0	1	2	3 ¤	g  row6¤	g  add-normal-slot¤	g  add-carriage-return-slot¤	g  add-full-line¤	g  initialize-playing-area¤	g  set-ace-low¤	g  make-standard-deck¤	g  shuffle-deck¤	 g  deal-cards-face-up-from-deck¤	!g  DECK¤	"g  append¤	#g  give-status-message¤	$g  new-game¤	%g  empty-slot?¤	&g  recalc-score¤	'g  
set-score!¤	(g  button-clicked¤	)g  move-n-cards!¤	*g  get-top-card¤	+g  remove-card¤	,g  sidle-up¤	-g  	do-action¤	.g  
droppable?¤	/g  button-released¤	0g  	get-value¤	1g  matches-in-rank¤	2g  get-suit¤	3g  matches-in-suit¤	4g  button-pressed¤	5g  playable-1?¤	6g  playable-3?¤	7g  	playable?¤	8g  button-double-clicked¤	9g  game-won¤	:g  get-hint¤	;g  game-continuable¤	<g  	hint-move¤	=g  	make-hint¤	>g  find-playable-move¤	?g  get-options¤	@g  apply-options¤	Ag  timeout¤	Bg  set-features¤	Cg  droppable-feature¤	Dg  
set-lambda¤C 5    h   {  ] 4	 >  "  G  RRRRRR h    ?  ] 4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  6  7      g  filenamef  accordion.scm
	
																			%			(			-			6			9			>			G			J			O			X	 		[	 		`	 		i	!		l	!		q	!		z	"		}	"	 	"	 	#	 	#	 	#	 	$	 	 
  g  nameg  add-full-line CR !"#  hH    ] 4>   "  G  4>   "  G  4>   "  G  4>   "  G  4>   "  G  4>   "  G  4>   "  G  4>   "  G  4>   "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  44	
5>  "  G  4>   "  G  			 C          g  filenamef  accordion.scm
	'
		)			*		#	+		3	,		C	/		S	0		c	1		s	2	 	3	 	4	 	4	 	4	 ¤	5	 §	5	 ¬	5	 µ	6	 ¸	6	 ½	6	 Æ	7	 É	7	 Î	7	 ×	8	 Ú	8	 ß	8	 è	9	 ë	9	 ð	9	 ù	:	 ü	:		:	
	<		<	%$	<	-	=	C	>	 %	D
  g  nameg  new-game C$R%&     h        ]4 5$  
4 5C
C              g  	last-slot
		  g  filenamef  accordion.scm
	B
		D				C			F				F			F				F	 		  g  nameg  recalc-score C&R'&     h   o   ] 4	356   g       g  filenamef  accordion.scm
	J
		K			K	 		
  g  nameg  give-status-message C#R    h   p   ]C    h       g  slot-id
		  g  filenamef  accordion.scm
	R
 		  g  nameg  button-clicked C(R%)*+,       h`     ] 	3$  N4 5$  "  4 5$  04  4 5 5$  4 5$   6CCCC           g  
first-slot
		[  g  filenamef  accordion.scm
	V
		W			W			X			X			X			W			Y		'	W		(	\		-	\		0	\	=	5	\	K	7	\	=	:	\	7	<	\		@	[		A	]		F	]		H	]		L	[		Q	^		S	^	 		[  g  nameg  sidle-up C,R+)%,#  hP   ó   ]4 5$  ?4 5$  .45$  "  45$  45$  6 CCCC  ë       g  end-slot
		N g  
start-slot		N g  	card-list			N  g  filenamef  accordion.scm
	e
		g	
		f			h	
		f			i		'	i	
	-	i	-	7	f		8	j	
	B	f		F	k	
 		N	  g  nameg  	do-action C-R.-  h    ¹   ]4 5$  
 6C   ±       g  
start-slot
		 g  	card-list		 g  end-slot			  g  filenamef  accordion.scm
	o
		p			p			q	 			  g  nameg  button-released C/R0*    h(   ¹   ] 
$  44 5545CC       ±       g  slot1
		! g  card		!  g  filenamef  accordion.scm
	v
		w		
	w			x			x			x			y			x	 			!	  g  nameg  matches-in-rank C1R2*    h(   ¾   ] 
$  44 5545CC       ¶       g  slot1
		! g  card		!  g  filenamef  accordion.scm
	~
				
			 		 		 		 		 	 			!	  g  nameg  matches-in-suit C3R% h   £   ]4 5$  C 
C           g  slot-id
		 g  	card-list		  g  filenamef  accordion.scm
 
	 		 		 		 			  g  nameg  button-pressed C4R56  h    ª   ]
4 5$  C 6  ¢       g  	from-slot
		 g  card		 g  t			  g  filenamef  accordion.scm
 
	 		 		 	 			  g  nameg  	playable? C7R31   h0   Ô   ]
 	$  "4 	5$  C 	6C Ì       g  from
		/ g  card		/ g  t			-  g  filenamef  accordion.scm
 
	 			 		 		 	#	 		 		) 	#	- 	 
		/	  g  nameg  playable-3? C6R31 h0   Ô   ]
 $  4 5$  C 6C      Ì       g  from
		* g  card		* g  t			(  g  filenamef  accordion.scm
 
	 	
	
 		  		  	!	  		  		$ ¡	!	( ¡	 
		*	  g  nameg  playable-1? C5R%6*-5   h`     ]4 5$  C4 4 55$   	 4 5 64 4 55$    4 5 6C         g  slot-id
		]  g  filenamef  accordion.scm
 ¦
	 §		 §		 ¨		 ¨	 	 ¨		! §		( ©		+ ©	4	4 ©	.	6 ©		7 «		< «	 	D «		H §		M ¬		P ¬	4	Y ¬	.	[ ¬	 		]  g  nameg  button-double-clicked C8R#9: h(      ] 4>   "  G  45 $  C6        y       g  filenamef  accordion.scm
 ´
	 µ		 ¶		 ¶		! ·	 		!
  g  nameg  game-continuable C;R%      h   y   ] 45$  4
5CC  q       g  filenamef  accordion.scm
 ½
	 ¾		 ¾		 ¿		 ¿	 		
  g  nameg  game-won C9R<      h      ] $    6C        g  possible-move
		  g  filenamef  accordion.scm
 Ã
	 Ä			 Ä		 Å		 Å	,	 Å	 		  g  nameg  	make-hint C=R%6*5>        hX     ]4 5$   C4 4 55$    	 C4 4 55$  	   C 6       	      g  
start-slot
		Q  g  filenamef  accordion.scm
 Ë
	 Ì		 Ì		 Í		 Ï		 Ï	$	  Ï		$ Ì		+ Ð		. Ð		0 Ò		5 Ò	$	= Ò		A Ì		F Ó		I Ó		O Ö	"	Q Ö	 		Q  g  nameg  find-playable-move C>R=>    h   g   ] 456    _       g  filenamef  accordion.scm
 Ü
	 Ý		 Ý	 		
  g  nameg  get-hint C:R%13    hX   9  ]45$  C $  "  	 $  45$  C6C       1      g  
start-slot
		Q g  	card-list		Q g  end-slot			Q g  t			- g  t		;	O  g  filenamef  accordion.scm
 à
	 â		 á		 ã		 ã		 ã			' ä		* ä		1 á		2 æ		9 æ	'	; æ		; æ			M ç	'	O ç	 		Q	  g  nameg  
droppable? C.R  h   X   ] C    P       g  filenamef  accordion.scm
 ì
 		
  g  nameg  get-options C?R   h   p   ]C    h       g  options
		  g  filenamef  accordion.scm
 î
 		  g  nameg  apply-options C@R   h   T   ] C    L       g  filenamef  accordion.scm
 ð
 		
  g  nameg  timeout CAR4BiCi>  "  G  Di$i4i/i(i8i;i9i:i?i@iAi.i6     s      g  filenamef  accordion.scm		
					
	 			#	
	%			(	
	*			-	
	/			2	
	4			7	
*	
+	'
	B
	J
#	R
À	V

	e
	o
	v
	~
Î 
­ 
Ç 
ß 
u ¦
4 ´
Ü ½
ª Ã
4 Ë
Â Ü
l à
Û ì
c î
Ï ð
Ð ò
 ô
 (	
   C6 