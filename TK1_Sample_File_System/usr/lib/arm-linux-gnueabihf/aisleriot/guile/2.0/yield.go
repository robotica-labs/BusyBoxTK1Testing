GOOF----LE-4-2.0_4      ] Q 4     h      ] g  guile¤	 ¤	g  process-use-modules¤	 ¤	 ¤	g  	aisleriot¤	g  	interface¤	 ¤		 ¤	
g  api¤	
 ¤	 ¤	g  initialize-playing-area¤	g  set-ace-low¤	g  make-standard-deck¤	g  shuffle-deck¤	g  add-normal-slot¤	g  DECK¤	g  HORIZPOS¤	e  2/3¤	g  add-extended-slot¤	g  right¤	g  add-carriage-return-slot¤	g  VERTPOS¤	e  0.5¤	g  add-blank-slot¤	g  deal-cards-face-up¤										
																				 ¤	g  give-status-message¤	g  new-game¤	g  set-statusbar-message¤	 g  get-stock-no-string¤	!g  string-append¤	"g  _¤	#f  Stock left:¤	$f   ¤	%g  number->string¤	&g  length¤	'g  	get-cards¤	(g  empty-slot?¤	)g  	get-value¤	*g  king¤	+g  
available?¤	,g  button-pressed¤	-g  get-top-card¤	.g  
droppable?¤	/g  add-to-score!¤	0g  remove-card¤	1g  reverse¤	2g  
set-cards!¤	3g  button-released¤	4 ¤	5	 ¤	6g  button-clicked¤	7g  button-double-clicked¤	8g  game-won¤	9g  get-hint¤	:g  game-continuable¤	;g  club¤	<f  Remove the king of clubs.¤	=g  diamond¤	>f  Remove the king of diamonds.¤	?g  heart¤	@f  Remove the king of hearts.¤	Ag  spade¤	Bf  Remove the king of spades.¤	Cg  hint-remove-king¤	Dg  
check-move¤	Eg  
hint-click¤	Fg  get-suit¤	Gg  	hint-move¤	Hf  Deal a card¤	Ig  	dealable?¤	Jg  check-waste¤	Kg  get-options¤	Lg  apply-options¤	Mg  timeout¤	Ng  set-features¤	Og  droppable-feature¤	Pg  
set-lambda¤C 5hð,  õ   ] 4	 >  "  G   hÈ  e  ] 4>   "  G  4>   "  G  4>   "  G  4>   "  G  4>  "  G  4>  "  G   4	
>  "  G  4>   "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>   "  G    4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>   "  G   4>   "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>   "  G    4>   "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>   "  G   4>   "  G  4>   "  G  4>  "  G  4>  "  G  4>  "  G  4>   "  G    4>   "  G  4>   "  G  4>  "  G  4>  "  G  4>   "  G   4>   "  G  4>   "  G  4>   "  G  4>  "  G  4
>  "  G  4>   "  G  		 C   ]      g  filenamef  	yield.scm
	
							#			3			C			U			X			]			i			j			l			m			p			w		 		 		 		 		 ¡	 	 ¤	 	 ©	 	 ²	!	 µ	!	 º	!	 Ã	"	 Æ	"	 Ë	"	 Ô	#	 ×	#	 Ü	#	 å	$	 è	$	 í	$	 ö	%	 ù	%	 þ	%		'		(		(		(	!	)	"	)	$	)	%	+	(	+	-	+	6	,	9	,	>	,	G	-	J	-	O	-	X	.	[	.	`	.	i	/	l	/	q	/	z	0	}	0		0		2		3		3	¡	3	¢	5	²	6	µ	6	º	6	Ã	7	Æ	7	Ë	7	Ô	8	×	8	Ü	8	å	9	è	9	í	9	ö	:	ù	:	þ	:		<		=		=		=	!	>	"	>	$	>	%	@	5	A	8	A	=	A	F	B	I	B	N	B	W	C	Z	C	_	C	h	D	k	D	p	D	y	F		G		G		G		I	 	J	°	K	³	K	¸	K	Á	L	Ä	L	É	L	Ò	M	Õ	M	Ú	M	ã	O	ö	P	÷	P	ù	P	ý	Q	þ	Q	 	Q		S		T	!	U	$	U	)	U	2	V	5	V	:	V	C	X	V	Y	W	Y	Y	Y	Z	[	j	\	z	]		^		^		^		`	 	`	¥	`	®	c	Ä	e	 	Å
  g  nameg  new-game CR         h   k   ] 45 6     c       g  filenamef  	yield.scm
	g
		h			h	 		
  g  nameg  give-status-message CR!"#$%&'        h    «   ] 45444
5556 £       g  filenamef  	yield.scm
	j
		k				k			k			k	"		l			l	!		l	)		l	!		l			k	 		
  g  nameg  get-stock-no-string C R()*+&    h8   Û   ]4 5$  C45$  C4 
5$  
45CCÓ       g  slot-id
		8 g  	card-list		8  g  filenamef  	yield.scm
	n
		o			o			p			p			p			p			o		!	q		,	o		-	r	
	5	r	 		8	  g  nameg  button-pressed C,R(    h  ¿  ]
 $  "   	$  C 
$  C 	$  	$  C	6 	$  	$  C	6 	$  	$  C	6 	$  ,4	5$  4	5$  	$  C	CCC 	$  	$  C	6 	$  	$  C	6 	$  ,4	5$  4	5$  	$  C	CCC 	$  ,4	5$  4	5$  	$  C	CCC 	$  	$  C	6 	$  	$  C	6 	$  ,4	5$  4	5$  	$  C	CCC 	$  ,4	5$  4	5$  	$  C	CCC 	$  ,4	5$  4	5$  	$  C	CCC 	$  	$  C	6 	$  	$  C	6 	$  ,4	5$  4	5$  	$  C	CCC 	$  ,4	5$  4	5$  	$  C	CCC 	$  ,4	5$  4	5$  	$  C	CCC 	$  ,4	5$  4	5$  	$  C	CCC 	
$  	$  C	6 		$  	$  C	6 	$  ,4	5$  4	5$  	$  C	CCC 	$  ,4	5$  4	5$  	$  C	CCC 	$  ,4	5$  4	5$  	$  C	CCC 	$  ,4	5$  4	5$  	$  C	CCC 	$  ,4	
5$  4	5$  	
$  C	CCC 	$  	
$  C	
6C·      g  slot-id
	 g  r-slot	 g  t			  g  filenamef  	yield.scm
	t
		u			u				v			u		%	x			)	u		0	z			4	u		9	{		=	{			E	|		J	}			N	u		S	~		W	~			_			d 			h	u		m 		q 			y 		~ 		 	u	  	  		  	  		  	 ¡ 		 ¨ 	 © 	 ³ 		 ·	u	 ¼ 	 À 		 È 	 Í 		 Ñ	u	 Ö 	 Ú 		 â 	 ç 		 ë	u	 ì 	 ö 		 ÷ 	 		 	
 		 	 	 		 	u	! 	+ 		, 	6 		; 	? 		F 	G 	Q 		U	u	Z 	^ 		f 	k 		o	u	t 	x 		 	 			u	 	 		  	 		¤ ¡	¨ 		¯ ¢	° ¢	º £		¾	u	¿ ¤	É ¤		Ê ¥	Ô ¤		Ù ¦	Ý ¤		ä §	å §	ï ¨		ó	u	ô ©	þ ©		ÿ ª		 ©		 «	 ©		 ¬	 ¬	$ ­		(	u	- ®	1 ®		9 ¯	> °		B	u	G ±	K ±		S ²	X ³		\	u	] ´	g ´		h µ	r ´		w ¶	{ ´		 ·	 ·	 ¸			u	 ¹	 ¹		 º	§ ¹		¬ »	° ¹		· ¼	¸ ¼	Â ½		Æ	u	Ç ¾	Ñ ¾		Ò ¿	Ü ¾		á À	å ¾		ì Á	í Á	÷ Â		û	u	ü Ã	 Ã		 Ä	 Ã		 Å	 Ã		! Æ	" Æ	, Ç		0	u	5 È	9 È		A É	F Ê		J	u	O Ë	S Ë		[ Ì	` Í		d	u	e Î	o Î		p Ï	z Î		 Ð	 Î		 Ñ	 Ñ	 Ò			u	 Ó	¤ Ó		¥ Ô	¯ Ó		´ Õ	¸ Ó		¿ Ö	À Ö	Ê ×		Î	u	Ï Ø	Ù Ø		Ú Ù	ä Ø		é Ú	í Ø		ô Û	õ Û	ÿ Ü			u	 Ý	 Ý		 Þ	 Ý		 ß	" Ý		) à	* à	4 á		8	u	9 â	C â		D ã	N â		S ä	W â		^ å	_ å	i æ		m	u	r ç	v ç		~ è	 Ú		  g  nameg  
available? C+R(+)-  h8     ]45$  C4 5$  	454455CC ý       g  
start-slot
		7 g  	card-list		7 g  end-slot			7  g  filenamef  	yield.scm
 ê
	 ë		 ë		 ì		 ë		 í		$ í		& í		' î		* î		2 î		3 í		4 í	 		7	  g  nameg  
droppable? C.R./0('12      h    Ë  ]4 5$  4	5$  z45$  m4	5$  "   $  "  $  ?4	5454 >  "  G  	44556CCCC      Ã      g  
start-slot
	  g  	card-list	  g  end-slot		  g  t		;	O g  new-contents		Z  g  moving-back		d	  g  filenamef  	yield.scm
 ð
	 ñ		 ñ		 ò		 ñ		 ó		' ñ		( ô		2 ô		; õ		; õ		L ö		S ô		T ÷		Z ÷		] ø	%	d ø	 	d ø		g ù		o ù		t ù	  ú	  ú	)  ú	$  ú	  ú	 	 	  g  nameg  button-released C3R(45+)-*0'12/    hè   9  ] 
$  045$  4
5$  C
64
5$  C
64 5$  C4 
5$  44 55$  x4	 5$  k $  W4	5$  "  C4
	5454 >  "  G  4	44555"  $  6CCCC      1      g  slot-id
	 â g  new-contents  Ì g  moving-back	  µ  g  filenamef  	yield.scm
 ý
	 þ		
 þ		 ÿ	
	 ÿ		 		 	
	%	%	'		(		1	
	8	%	:		;		E		H		S		T		W		_		b		f		g		q		u			y			z
	 
	 	) 	 	/ 	* 	 	 ¥	' ª	 ¶	 »	% ¾	3 Å	. Ç	% É	 Õ	 Ú	 -	 â  g  nameg  button-clicked C6R6      h   }   ] 6u       g  slot-id
		  g  filenamef  	yield.scm

		 		  g  nameg  button-double-clicked C7R89      h(   }   ] 4>   "  G  45 $  C6        u       g  filenamef  	yield.scm

							!	 		!
  g  nameg  game-continuable C:R(  hp   ê   ] 4
5$  `45$  T4	5$  G4	5$  :4	5$  -4	5$   4	5$  4	5$  		6CCCCCCCC  â       g  filenamef  	yield.scm

											!		"		,		-		7		8 		B		C!		M		N"		X		^#	 		n
  g  nameg  game-won C8R;"<=>?@AB 
    h@   Ø   ] &  6 &  6 &  6 &  	6C    Ð       g  suit
		<  g  filenamef  	yield.scm
%
	
&		&		&		&		'		'		&&		*(		,(		4&		8)		:)	 		<  g  nameg  hint-remove-king CCR(+D*)-ECFG     hð   a  ]4 5$  "  	4 
5$   	$   	 6C44 55$   44	4 555645$  "  64
5$  "  	44 554455$  )	$  	 6 	$   	 6C
 6      Y      g  slot1
	 ê g  slot2	 ê g  t				" g  t		n ´ g  t	  ±  g  filenamef  	yield.scm
+
	,	
		,		-		-	
	&,		+.	
	/.		4/		9/	"	;/	
	@1		C1		K1		L1	
	P1		U2		X2	.	[2	8	c2	.	e2		g2	
	h3		n3		|4	 4	 3	 5	  5	+ 5	   6	  £6	+ «6	  ¬5	 ­5	 ®5	 ¸3	
 ½7	 Á7	 È8	$ Ê8	 Ï9	 Ó9	 Ø:	" Ý:	. ß:	 ê<	 0	 ê	  g  nameg  
check-move CDR("H  h       ] 4
5$  C
45 C             g  filenamef  	yield.scm
>
	?		?		@		@		@		@	 		
  g  nameg  	dealable? CIR(&')-G     hP   ö   ] 4	5$  C44	55$  -	44	5544	55$  			6CCî       g  filenamef  	yield.scm
C
	D		D		E	
	E		E	
	E		!D		$F		'F		/F		0G		3G	!	:G		=G		>F		?F		CD		LH	 		P
  g  nameg  check-waste CJRDJI     h0      ]4	5  $   C45   $   C6               g  t
	
	) g  t
		)  g  filenamef  	yield.scm
J
	K		
K		L		K		)M	 		)
  g  nameg  get-hint C9R       h   T   ] C    L       g  filenamef  	yield.scm
O
 		
  g  nameg  get-options CKR       h   l   ]C    d       g  options
		  g  filenamef  	yield.scm
R
 		  g  nameg  apply-options CLR       h   P   ] C    H       g  filenamef  	yield.scm
U
 		
  g  nameg  timeout CMR4NiOi>  "  G  Pii,i3i6i7i:i8i9iKiLiMi.i6 í       g  filenamef  	yield.scm		
	x	

	g

þ	j
.	n
	t
Ø ê
f ð
 ´ ý
!P
"
#}
$»%
(4+
(ó>
*YC
+?J
+¯O
,7R
,£U
,¤X
,ïZ
 	,ï
   C6 