GOOF----LE-4-2.0L#      ] Z 4    ha      ] g  guile¤	 ¤	g  process-use-modules¤	 ¤	 ¤	g  	aisleriot¤	g  	interface¤	 ¤		 ¤	
g  api¤	
 ¤	 ¤	g  initialize-playing-area¤	g  set-ace-low¤	g  make-standard-deck¤	g  shuffle-deck¤	g  add-normal-slot¤	g  DECK¤	g  add-blank-slot¤	g  add-carriage-return-slot¤	g  deal-cards-face-up¤										 ¤	g  check-kings¤	g  add-to-score!¤	g  give-status-message¤	g  new-game¤	g  set-statusbar-message¤	g  string-append¤	g  get-stock-no-string¤	f     ¤	g  get-redeals-string¤	 g  _¤	!f  Stock left:¤	"f   ¤	#g  number->string¤	$g  length¤	%g  	get-cards¤	&f  Redeals left:¤	'g  FLIP-COUNTER¤	(g  	get-value¤	)g  get-top-card¤	*g  king¤	+g  
set-cards!¤	,g  reverse¤	-g  	make-card¤	.g  get-suit¤	/g  remove-card¤	0g  empty-slot?¤	1g  button-pressed¤	2g  
droppable?¤	3g  move-n-cards!¤	4g  
deal-cards¤	5g  button-released¤	6g  
flippable?¤	7g  	dealable?¤	8g  
flip-stock¤	9g  do-deal-next-cards¤	:g  button-clicked¤	;	 ¤	<g  button-double-clicked¤	=g  
check-move¤	>g  game-won¤	?g  	game-over¤	@g  ace¤	Af  You are searching for an ace.¤	Bf  You are searching for a two.¤	Cf  You are searching for a three.¤	Df  You are searching for a four.¤	Ef  You are searching for a five.¤	Ff  You are searching for a six.¤	Gf  You are searching for a seven.¤	Hf  You are searching for an eight.¤	If  You are searching for a nine.¤	Jf  You are searching for a ten.¤	Kg  jack¤	Lf  You are searching for a jack.¤	Mg  queen¤	Nf  You are searching for a queen.¤	Of  You are searching for a king.¤	Pf  Unknown value¤	Qg  get-value-hint¤	Rg  get-hint¤	Sg  get-options¤	Tg  apply-options¤	Ug  timeout¤	Vg  set-features¤	Wg  droppable-feature¤	Xg  dealable-feature¤	Yg  
set-lambda¤C 5  hp  ú   ] 4	 >  "  G   hà    ] 4>   "  G  4>   "  G  4>   "  G  4>   "  G  4>  "  G  4>  "  G  4>   "  G  4>  "  G  4>  "  G  4>  "  G  4>   "  G  4>   "  G  4>   "  G  4>   "  G  4>  "  G  4>   "  G  4>  "  G  4>   "  G  4>   "  G  4>   "  G  4>   "  G  4>  "  G  4>  "  G  4>  "  G  4	

>  "  G  4
>  "  G  4>  "  G  4>   "  G  		 C           g  filenamef  doublets.scm
	
							#			3			C			U			X			]			f			v			y			~		 		 		 		 		 		  		 ©	 	 ¹	"	 É	#	 Ù	$	 é	%	 ì	%	 ñ	%	 ú	&	
	'		'		'		(	+	*	;	+	K	,	[	-	^	-	c	-	l	.	o	.	t	.	}	/		/		/		1		1		1	¡	3	¥	3	ª	3	³	5	Ä	7	Ú	9	 5	Û
  g  nameg  new-game CR    h      ] 445 45 56        g  filenamef  doublets.scm
	<
		=			=	(		>	(		?	(		=			=	 		
  g  nameg  give-status-message CR !"#$%     h    ®   ] 45444
5556 ¦       g  filenamef  doublets.scm
	A
		B				B			B			B	"		C			C	!		C	)		C	!		C			B	 		
  g  nameg  get-stock-no-string CR &"#'       h       ] 454	56              g  filenamef  doublets.scm
	E
		F				F			F			F	$		G			G	!		G			F	 
		
  g  nameg  get-redeals-string CR()*%+,-./$    h¨   Ì  ]	44 55$  x4
54
4444 5544 555455>  "  G  4	 >  "  G  4

  >  "  G   64 5$   6C    Ä      g  	slot-list
	 ¤ g  new-deck		^  g  filenamef  doublets.scm
	I
		J				J			J	"		J			J				J			J			L			L			M		#	M		&	M	+	)	M	6	,	M	A	1	M	O	3	M	A	5	M	6	6	N	6	9	N	@	>	N	N	@	N	@	B	N	6	D	M	+	E	O	+	L	M	%	N	M		S	M		_	P		d	P		i	P		r	Q		x	Q	$	{	Q	4 	Q	 	R	 	S	 	S	
 	S	  	T	 ¢	T	
 *	 ¤  g  nameg  check-kings CR0(*       h8   Þ   ]4 5$  C 
$  C 	$  C45C       Ö       g  slot-id
		1 g  	card-list		1  g  filenamef  doublets.scm
	V
		W			W			X			W			Y		"	W		%	Z		*	Z		,	Z		/	Z		0	Z	 		1	  g  nameg  button-pressed C1R()       h0   ó   ]	$  45	4455	CC       ë       g  
start-slot
		) g  	card-list		) g  end-slot			)  g  filenamef  doublets.scm
	\
		]			]			^	
		^			^	
		_			_	"	!	_		"	_		%	_	
	&	^	 		)	  g  nameg  
droppable? C2R2304  hh      ]4 5$  Q4>  "  G  4 >  "  G  45$  4
5$  C
  6  6C    ø       g  
start-slot
		d g  	card-list		d g  end-slot			d  g  filenamef  doublets.scm
	a
		b			b			d			#	e			9	f		B	f			C	h		L	f			V	i	7	X	i		`	g	/	b	g	 		d	  g  nameg  button-released C5R6       h   \   ] 
	6      T       g  filenamef  doublets.scm
	l
	
	m	 		

  g  nameg  	dealable? C7R8   h   e   ] 
	6      ]       g  filenamef  doublets.scm
	o
	
	p	 		

  g  nameg  do-deal-next-cards C9R9  h      ] 
$  6 C       g  slot-id
		  g  filenamef  doublets.scm
	r
		s		
	s			t	 		  g  nameg  button-clicked C:R0()*4; 	 hÈ   ­  ]4 5$  "  R 
$  "  E 	$  "  744 55$  "  44 55	44	55	$  ]4>  "  G  4 >  "  G  45$  $4
5$  "   $  

  6C  6C ¥      g  slot
	 Ç  g  filenamef  doublets.scm
	v
		w			w			x			w		$	y		(	w		.	z		1	z		9	z		<	z		@	w		F	{		I	{		Q	{		T	|		W	|	&	_	|		`	|		c	|		d	{		h	w		i	~		z		 		 		  	  	  	 ¡ 	 ª 	 « 	 ¯ 	 · 	0 ¹ 	 Ã 	( Å 	 &	 Ç  g  nameg  button-double-clicked C<R0()=    hP   ù   ]4 5$  "  44 55	44	55	$  C 		$   6C       ñ       g  slot
		I  g  filenamef  doublets.scm
 
	 		 		 		 		 		! 		$ 	&	, 		- 		0 		1 		5 		< 	
	@ 		E 		G 	
 		I  g  nameg  
check-move C=R>'0=      hH   Ä   ]4>   "  G  45 $  C	  $   C4
5  $   C6    ¼       g  t
	"	D g  t
	4	D  g  filenamef  doublets.scm
 
	 		 		 		" 		" 		. 		4 		4 		D 	 		D
  g  nameg  	game-over C?R$% h   x   ] 	044	55C      p       g  filenamef  doublets.scm
 
	 		 		 		 	 		
  g  nameg  game-won C>R@ ABCDEFGHIJKLMN*OP   hÀ   ð  ] &  6 	&  6 	&  6 	&  6 	&  6 	&  6 	&  	6 	&  
6 		&  6 	
&  6 &  6 &  6 &  66  è      g  value
	 ¾  g  filenamef  doublets.scm
 
	
 		 		 		 		 		 		& 		* 		, 		4 		8 		: 		B 		F  		H  		P 		T ¡		V ¡		^ 		b ¢		d ¢		l 		p £		r £		z 		~ ¤	  ¤	  	  ¥	  ¥	  	  ¦	  ¦	 ¤ 	 ¨ §	 ª §	 ² 	 ¶ ¨	 ¸ ¨	 ¼ ©	 ¾ ©	 +	 ¾  g  nameg  get-value-hint CQR()Q   h(   ±   ]	44	55	 
4 5 C      ©       g  wanted
		"  g  filenamef  doublets.scm
 «
	 ¬		 ¬	(	 ¬		 ¬		 ¬		 ¬		 ®		! ®	 
		"
  g  nameg  get-hint CRR  h   W   ] C    O       g  filenamef  doublets.scm
 °
 		
  g  nameg  get-options CSR    h   o   ]C    g       g  options
		  g  filenamef  doublets.scm
 ²
 		  g  nameg  apply-options CTR    h   S   ] C    K       g  filenamef  doublets.scm
 ´
 		
  g  nameg  timeout CUR4ViWiXi>  "  G  Yii1i5i:i<i?i>iRiSiTiUi2i7i6      ò       g  filenamef  doublets.scm		
6	
ù	<
é	A
È	E
	g	I

	V
Ö	\
[	a
ß	l
h	o
	r
¨	v
 
7 
Û 
Ã 
´ «
" °
ª ²
 ´
 ¶
j ¸
 	j
   C6 