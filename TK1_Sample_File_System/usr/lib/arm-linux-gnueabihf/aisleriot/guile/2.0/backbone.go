GOOF----LE-4-2.08      ] f 4        h      ] g  guile¤	 ¤	g  process-use-modules¤	 ¤	 ¤	g  	aisleriot¤	g  	interface¤	 ¤		 ¤	
g  api¤	
 ¤	 ¤	
						 ¤	g  tableau¤									
	 ¤	g  
foundation¤	g  stock¤	g  waste¤																 	!	"	#	$ ¤	g  reserve¤	g  reserve?¤	g  member¤	g  	playable?¤	g  obscured¤	g  padded-list¤	g  add-normal-slot¤	g  HORIZPOS¤	g  VERTPOS¤	e  1/3¤	g  	list-tail¤	g  make-backbone¤	 g  deal-cards-face-up¤	!g  "deal-cards-face-up-to-reserve-from¤	"g  deal-cards-face-up-to-reserve¤	#g  initialize-playing-area¤	$g  set-ace-low¤	%g  make-standard-double-deck¤	&g  shuffle-deck¤	'g  add-blank-slot¤	(g  add-carriage-return-slot¤	)g  DECK¤	*e  1.5¤	+g  give-status-message¤	,g  new-game¤	-g  set-statusbar-message¤	.g  string-append¤	/g  get-stock-no-string¤	0f     ¤	1g  get-redeals-string¤	2g  _¤	3f  Redeals left:¤	4f   ¤	5g  number->string¤	6g  FLIP-COUNTER¤	7f  Stock left:¤	8g  length¤	9g  	get-cards¤	:g  empty-slot?¤	;g  empty-slots?¤	<g  list-ref¤	=g  is-playable?¤	>g  	get-value¤	?g  get-suit¤	@g  get-top-card¤	Ag  is-legal-move?¤	Bg  button-pressed¤	Cg  move-n-cards!¤	Dg  add-to-score!¤	Eg  complete-transaction¤	Fg  button-released¤	Gg  
flippable?¤	Hg  	dealable?¤	Ig  
flip-stock¤	Jg  do-deal-next-cards¤	Kg  button-clicked¤	Lg  remove-card¤	Mg  move-if-possible¤	Ng  button-double-clicked¤	Og  non-empty-piles-helper¤	Pg  non-empty-piles¤	Qg  empty-piles¤	Rg  get-legal-move-from-source¤	Sg  get-legal-move¤	Tg  "can-send-tableau-card-to-any-slot?¤	Ug  useful-tableau-stack?¤	Vg  get-move-from-tableau¤	Wg  append¤	Xg  	hint-move¤	Yf  Deal a new card from the deck¤	Zf  Try rearranging the cards¤	[g  get-hint¤	\g  full-piles?¤	]g  game-won¤	^g  	game-over¤	_g  get-options¤	`g  apply-options¤	ag  timeout¤	bg  set-features¤	cg  droppable-feature¤	dg  dealable-feature¤	eg  
set-lambda¤C 5       h .  Ù  ] 4	 >  "  G  RR	R	RR    h   n   ] 	Cf       g  slot
		  g  filenamef  backbone.scm
	
			 		  g  nameg  reserve? CR     h0   «   ]	 	$  C4 5$  C 6    £       g  slot
		, g  t		, g  t		,  g  filenamef  backbone.scm
	
													,		 		,  g  nameg  	playable? CRR        h   ¨   ] 
$  C 6        g  n
		 g  tail		  g  filenamef  backbone.scm
	#
		$		
	$			&			&	!		&			&	 			  g  nameg  padded-list CR 	hH   ð   ]4>  "  G    4 5  	$   C 6    è       g  n
		D  g  filenamef  backbone.scm
	(
		)			)			)			)			*			*			+			+		!	+		"	,		-	,	)	0	,	#	1	,		6	-		:	-		B	/		D	/	 		D  g  nameg  make-backbone CR !   h0   µ   ]4  >  "  G   	$$  C 6     ­       g  n
		+  g  filenamef  backbone.scm
	1
		2			2			2			3		"	3		)	5	*	+	5	 			+  g  nameg  "deal-cards-face-up-to-reserve-from C!R!  h   p   ] 	6h       g  filenamef  backbone.scm
	7
		8	 		
  g  nameg  deal-cards-face-up-to-reserve C"R#$%&'() *"+       hð  -  ] 4>   "  G  4>   "  G  4	)5 4>   "  G  4>   "  G  4>  "  G  4	>   "  G  4	>   "  G  4>  "  G  4
>  "  G  4
>  "  G  4
>  "  G  4
>  "  G  4>   "  G  4>  "  G  4	>   "  G  4	>   "  G  4>  "  G  4
>  "  G  4
>  "  G  4
>  "  G  4
>  "  G  4>   "  G  4>  "  G  4	>   "  G  4	>   "  G  4>  "  G  4	>   "  G  4>  "  G  4>  "  G  4>   "  G  4>  "  G  4	>   "  G  4	>   "  G  4>  "  G  4>  "  G   
 44	55	$ 	 
 4	>  "  G   4>  "  G  4>   "  G  4>   "  G  		 C       %      g  filenamef  backbone.scm
	:
		;			<		#	=		(	=	!	*	=		,	=		-	?		=	@		M	B		P	B		R	B		W	B		`	C		p	D	 	E	 	E	 	E	 	E	 	F	 	F	 	F	 	F	 ¦	G	 ©	G	 «	G	 °	G	 ¹	H	 ¼	H	 ¾	H	 Ã	H	 Ì	I	 Ï	I	 Ñ	I	 Ö	I	 ß	K	 ï	L	 ò	L	 ô	L	 ù	L		M		N	"	O	%	O	'	O	,	O	5	P	8	P	:	P	?	P	H	Q	K	Q	M	Q	R	Q	[	R	^	R	`	R	e	R	n	S	q	S	s	S	x	S		U		V		V		V		V	¤	W	´	X	Ä	Y	Ç	Y	É	Y	Î	Y	×	Z	ç	[	í	[	ò	[	û	\	þ	\	 	\		\		^		_	!	_	#	_	(	_	1	`	A	a	Q	b	T	b	V	b	[	b	d	d	z	f	}	g	~	h		h	 	h		h	4	h		j		k		l	¬	n	®	n	¯	o	²	o	´	o	¹	o	Â	q	Ò	s	è	u	 q	é
  g  nameg  new-game C,R-./01  h      ] 445 45 56        g  filenamef  backbone.scm
	x
		y			y	(		z	(		{	(		y			y	 		
  g  nameg  give-status-message C+R.23456       h      ] 45456       g  filenamef  backbone.scm
	}
		~				~			~			~	$						!					~	 
		
  g  nameg  get-redeals-string C1R.274589 	   h    ¹   ] 454445556±       g  filenamef  backbone.scm
 
	 			 		 		 	"	 		 	!	 	)	 	!	 		  	 		 
  g  nameg  get-stock-no-string C/R:;    h    ´   ] &  C4 5$   6C¬       g  slots
		   g  filenamef  backbone.scm
 
	 			 		 			 		 			 		 		 		 
		   g  nameg  empty-slots? C;R;<       hP   Ð   ]4 5$  "   $  "  4 5$  4 56C      È       g  slot
		J g  t		7 g  t			4  g  filenamef  backbone.scm
 
	 			 		 			 		+ 			; 		> 		H 	 
		J  g  nameg  is-playable? C=R=:>?@ 
      hè   ©  ] $  C&  C4 5$  ¿45$  Q45$  45"  4454455$  454455"  "  $  C45$  F45$  	4	 5C454455$  454455CCCC      ¡      g  
start-slot
	 â g  	card-list	 â g  end-slot		 â g  t	  à  g  filenamef  backbone.scm
 
	 		 		 		 		 		! 		" 		. 		/ 		9 		: 		? 	"	A 		C 		H 		M 	#	O 		P 		S 	#	[ 		\ 		` 		a 		f 	$	h 		i 		l 	'	t 		u 		v 	  	  	  	   	 £  	 ¤ ¡	 « ¡	 ­ £	 ² £	# ´ £	 µ ¤	 ¸ ¤	# À ¤	 Á £	 Å ¢	 Æ ¥	 Ë ¥	$ Í ¥	 Î ¦	 Ñ ¦	' Ù ¦	 Ú ¦	 Û ¥	 7	 â	  g  nameg  is-legal-move? CAR=      h      ] 6       g  slot-id
		 g  	card-list		  g  filenamef  backbone.scm
 ¨
	 ©	 			  g  nameg  button-pressed CBRCD h`   å   ]4 >  "  G  4 5$  4	ÿ>  "  G  "   45$  4>  "  G  "   C Ý       g  
start-slot
		_ g  	card-list		_ g  end-slot			_  g  filenamef  backbone.scm
 «
	 ¬		 ­		% ­		& ®		< ¯		H ¯		I °	 			_	  g  nameg  complete-transaction CERAE        h(   Î   ] $  C4 5$  
 6CÆ       g  
start-slot
		( g  	card-list		( g  end-slot			(  g  filenamef  backbone.scm
 ³
	 ´		 ´		 µ		 ´		& ¶	 		(	  g  nameg  button-released CFRG     h   ^   ] 6     V       g  filenamef  backbone.scm
 ¸
	 ¹	 		
  g  nameg  	dealable? CHRI     h   g   ] 6    _       g  filenamef  backbone.scm
 »
	 ¼	 		
  g  nameg  do-deal-next-cards CJRI    h      ] $  
6C        g  
start-slot
		  g  filenamef  backbone.scm
 ¾
	 ¿		 ¿		 À	 		  g  nameg  button-clicked CKR:@ALEM       h`   @  ]
4 5$  C4 54  5$   4 >  "  G    6&  C 6  8      g  
start-slot
		^ g  	end-slots		^ g  card			^  g  filenamef  backbone.scm
 Â
	 Ã		 Ã		 Ä		 Ä		 Å		" Å	(	% Å	4	' Å		+ Å			, Æ		F Ç	3	I Ç	?	K Ç		N È		O È	&	S È		\ Ê	.	^ Ê	 		^	  g  nameg  move-if-possible CMRM     h      ] 6      {       g  
start-slot
		
  g  filenamef  backbone.scm
 Ì
	
 Í	 		
  g  nameg  button-double-clicked CNRO:  h0   ä   ] &  C 4 5$  "   6     Ü       g  piles
		+ g  result		+  g  filenamef  backbone.scm
 Ï
	 Ð			 Ð		 Ó		 Ô		 Ô		 Ô		 Ô		& Ö		) Ö		+ Ò	 		+	  g  nameg  non-empty-piles-helper CORO   h      ] 6       y       g  piles
			  g  filenamef  backbone.scm
 Ø
	 Ù	 		 Ù	 			  g  nameg  non-empty-piles CPR:      h@     ]"  -&  C45$  
"  "ÿÿÓ "ÿÿÈ  þ       g  piles
		> g  piles		3 g  result			3  g  filenamef  backbone.scm
 Û
	 Ü			 Ý		 Ý		 à	
	 á		 á		 á		 á	
	" â		% â		3 ß		6 ä		> ä	 		>  g  nameg  empty-piles CQR:A9R   hH     ]&  C4 5$  "  4 4 5 5$  	  C 6        g  source
		F g  targets		F  g  filenamef  backbone.scm
 æ
	 ç			 ç		 é		 é	
	 ê		! ê	1	( ê	,	+ ê	&	. ê	F	0 ê		4 é		9 ë		< ë	
	D ì	-	F ì	
 		F	  g  nameg  get-legal-move-from-source CRRRS      h0   Ü   ]
 &  C4 5$  C 6       Ô       g  sources
		) g  targets		) g  t			)  g  filenamef  backbone.scm
 î
	 ï			 ï		 ñ	
	 ñ	&	 ñ	
	 ñ		% ò		) ò	
 
		)	  g  nameg  get-legal-move CSRAT       h0   ý   ]
&  C4  5$  C 6 õ       g  card
		/ g  slots		/ g  t			/  g  filenamef  backbone.scm
 ô
	 õ			 õ		 ÷	
	 ÷		 ÷	(	 ÷	4	 ÷	
	 ÷		- ø	3	/ ø	
 		/	  g  nameg  "can-send-tableau-card-to-any-slot? CTRTPU    h(   Æ   ] &  C4 455$   6C ¾       g  cards
		'  g  filenamef  backbone.scm
 ú
	 û			 û		 ý		 ý	/	 ý	;	 ý		 ý		# þ	"	% þ	 		'  g  nameg  useful-tableau-stack? CURRPVU9     h`   _  ]	 &  C4 5$  C4 455"   6$  44 55$  C"ÿÿÚ"ÿÿÖ W      g  sources
		_ g  move		_ g  move	/	_  g  filenamef  backbone.scm
 
									.							!		&	7	'	E	/		/		:
	/	<
		<		C		F	9	K	D	M	9	N	4	P		T	 		_  g  nameg  get-move-from-tableau CVRSWPVRQX:62YZ        hð   7  ]44 5445 55  $   "  &45  $   "  4	4
55    $  4  5"     $   C45  $  "  $  45"   $  
45 "    $   C45$  
45 CC    /      g  t
	&	] g  t
	;	Z g  move
	]	z g  t
	z ì g  t
  ² g  t
 Æ ì  g  filenamef  backbone.scm

						%						(	"	B	$		&		&		5		;		J		O		W		]		e		f		k		o	$	r		z	 	 	 		 	 ¢	 £	 ª	 ¶	 ¸	 ¼	 ¾	 Á		 Æ	 Ò		 Þ	 à	 ä	 æ	 é		 +	 ì
  g  nameg  get-hint C[R89\    h0   î   ]		44 55$   $  C 6C æ       g  piles
		/ g  t		-  g  filenamef  backbone.scm
!
	"		"		"	 	"		"		"		"		#		#		#		#		+$		-$	 		/  g  nameg  full-piles? C\R\       h   ]   ] 6U       g  filenamef  backbone.scm
&
	'	 		
  g  nameg  game-won C]R+[        h   g   ] 4>   "  G  6   _       g  filenamef  backbone.scm
)
	*		,	 		
  g  nameg  	game-over C^R    h   `   ] C    X       g  filenamef  backbone.scm
.
	/	 		
  g  nameg  get-options C_R   h   o   ]C    g       g  options
		  g  filenamef  backbone.scm
1
 		  g  nameg  apply-options C`R    h   S   ] C    K       g  filenamef  backbone.scm
3
 		
  g  nameg  timeout CaR4bicidi>  "  G  ei,iBiFiKiNi^i]i[i_i`iaiAiHi6      Ñ      g  filenamef  backbone.scm		
					
	 			#	
	(	
	-	
	/			2	
 ¹	
®	
¯	!	²	!
	#
ã	(
à	1
k	7
Ð	:
	x
h	}
d 
O 
 
L 
û ¨
X «
i ³
ñ ¸
 »
A ¾
 Â
® Ì
× Ï
| Ø
Ù Û
!R æ
"w î
#À ô
$É ú
&ª 
*

+A!
+À&
,Z)
,Ó.
-Z1
-Æ3
-Ç5
.7
 2	.
   C6 