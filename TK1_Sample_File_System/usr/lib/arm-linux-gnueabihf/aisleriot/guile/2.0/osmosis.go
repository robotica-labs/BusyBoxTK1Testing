GOOF----LE-4-2.0û*      ] W 4       h      ] g  guile¤	 ¤	g  process-use-modules¤	 ¤	 ¤	g  	aisleriot¤	g  	interface¤	 ¤		 ¤	
g  api¤	
 ¤	 ¤	g  ice-9¤	g  format¤	 ¤	 ¤	g  
deal-three¤	g  initialize-playing-area¤	g  set-ace-low¤	g  make-standard-deck¤	g  shuffle-deck¤	g  add-extended-slot¤	g  right¤	g  reserve¤	g  add-blank-slot¤	g  
foundation¤	g  add-carriage-return-slot¤	g  add-normal-slot¤	g  DECK¤	g  stock¤	g  add-partially-extended-slot¤	 g  waste¤	!g  initial-deal¤	"g  give-status-message¤	#g  add-to-score!¤	$g  new-game¤	%g  
deal-cards¤	&
			
			
			 ¤	'g  deal-cards-face-up¤	(
			 ¤	)g  set-statusbar-message¤	*g  string-append¤	+g  get-stock-no-string¤	,f     ¤	-g  get-redeals-string¤	.g  _¤	/f  Stock left: ~a¤	0g  number->string¤	1g  length¤	2g  	get-cards¤	3f   ¤	4f  Redeals left: ~a¤	5g  FLIP-COUNTER¤	6g  empty-slot?¤	7g  button-pressed¤	8g  move-n-cards!¤	9g  make-visible-top-card¤	:g  complete-transaction¤	;g  	get-value¤	<g  find-card-val-in-list?¤	=g  reverse¤	>g  get-suit¤	?g  get-top-card¤	@g  
droppable?¤	Ag  button-released¤	Bg  
flip-stock¤	Cg  button-clicked¤	Dg  remove-card¤	Eg  check-to-move¤	Fg  button-double-clicked¤	Gg  	hint-move¤	Hg  
placeable?¤	Ig  get-valid-move¤	J
					 ¤	Kg  game-continuable¤	Lg  game-won¤	Mf  Deal new cards from the deck¤	Nf  Deal a new card from the deck¤	Og  get-hint¤	Pf  Three card deals¤	Qg  get-options¤	Rg  apply-options¤	Sg  timeout¤	Tg  set-features¤	Ug  droppable-feature¤	Vg  
set-lambda¤C 5       hà"  	  ] 4	 >  "  G  R !"#       hð  a  ] 4>   "  G  4>   "  G  4>   "  G  4>   "  G  4>  "  G  4>   "  G  4	>  "  G  4
>   "  G  4>  "  G  4>   "  G  4	>  "  G  4
>   "  G  4>  "  G  4>   "  G  4	>  "  G  4
>   "  G  4>  "  G  4>   "  G  4	>  "  G  4
>   "  G  4>  "  G  $  4	>  "  G  "  4>  "  G  4>   "  G  4>   "  G  4>  "  G  		 C     Y      g  filenamef  osmosis.scm
	
							#			3			C			F			J			O			X			h			k			o			t			}		 		 		 		 		 ¢		 ²	 	 µ	 	 ¹	 	 ¾	 	 Ç	!	 ×	"	 Ú	"	 Þ	"	 ã	"	 ì	#	 ü	$	 ÿ	$		$		$		%	!	&	$	&	(	&	-	&	6	'	F	(	I	(	M	(	R	(	[	)	k	*	q	*	v	*		,		-		-	!	-	-	-	 	.	£	.	¥	.	ª	.	³	1	Ã	3	Ó	5	ê	6	 >	ë
  g  nameg  new-game C$R%&'(        h    ~   ] 4	>  "  G  	6  v       g  filenamef  osmosis.scm
	8
		9				9			9			:			:	 		
  g  nameg  initial-deal C!R)*+,- h      ] 445 45 56        g  filenamef  osmosis.scm
	=
		>			>	(		?	(		@	(		>			>	 		
  g  nameg  give-status-message C"R./012        h    ¥   ] 45444	5556        g  filenamef  osmosis.scm
	B
		C		
	C			C			C	!		C	1		C	9		C	1		C	!		C	 		
  g  nameg  get-stock-no-string C+R3.405      h(   ¤   ] $  C454	56              g  filenamef  osmosis.scm
	F
		G		
	H			I			I			I			I	%		I	5		I	%	!	I	 		!
  g  nameg  get-redeals-string C-R61 hh   1  ]
4 5$  C45$  E 
$  C 	$  C 	$  C 	$  C 		CC     )      g  slot-id
		c g  	card-list		c g  t		 	a g  t		0	a g  t		@	a g  t		P	a  g  filenamef  osmosis.scm
	M
		N			N			O	
		O			N		 	P		 	P		0	Q		0	P		@	R		@	P		P	S		P	P		`	T	 		c	  g  nameg  button-pressed C7R8#69        hP   Ì   ]4 >  "  G  4>  "  G  4 5$  "  4 >  "  G  C    Ä       g  
start-slot
		L g  	card-list		L g  end-slot			L  g  filenamef  osmosis.scm
	V
		W			X		*	Y		4	Y		9	Z	 		L	  g  nameg  complete-transaction C:R;< h0   Ö   ]
 (  C4 5$  C 6       Î       g  cards
		) g  value		) g  t			)  g  filenamef  osmosis.scm
	]
		^			_			_			_			_			_		%	`	#	)	`	 
		)	  g  nameg  find-card-val-in-list? C<R6;=2>?<     hð   v  ]H J$  CJ$  "  /J	$  "  J	$  "  J	$  4J5$  K44455545$  +"  #"  4J	5$  J	K"ÿÿçC"ÿÿá"ÿÿÙC44J5545$  %J$  C4J	5456CCn      g  
start-slot
	 ð g  	card-list	 ð g  end-slot		 ð g  t			T g  t		'	Q g  t		9	N g  t	 Ë ì  g  filenamef  osmosis.scm
	b
		c			c			d			d		'	e		'	d		9	f		9	d		K	g		X	c		Y	h		c	h		d	i		g	i	#	j	i	,	q	i	#	r	i		t	i		u	j		z	j		|	j		}	i	 	i	 	k	 	k	 	k	$ 	k	 	k	 	l	& 	l	 ¢	k	 ¯	m	 ²	m	 º	m	 »	n	 À	n	 Â	n	 Ã	m	 Ç	m	 Ë	o	 Ë	o	 Ù	p	, à	p	7 â	p	, ã	q	, è	q	7 ê	q	, ì	p	 2	 ð	  g  nameg  
droppable? C@R@:       h    ·   ]4 5$  
 6C   ¯       g  
start-slot
		 g  	card-list		 g  end-slot			  g  filenamef  osmosis.scm
	s
		t			t			u	 			  g  nameg  button-released CARB      h0      ] 	$  #			$  	ÿ"  	$  	"  6C       g  slot-id
		0  g  filenamef  osmosis.scm
	w
		x			x			y		%	y	,	.	y	 		0  g  nameg  button-clicked CCR;D:E   hP     ](  C4545&  4 >  "  G    6 6             g  	orig-slot
		I g  end-slot		I g  
above-list			I g  top-card			I  g  filenamef  osmosis.scm
	{
		|			}			~			~			~			}		 		8 	,	< 		E 	,	I 	
 		I	  g  nameg  check-to-move CER6?>2D:;=E 
      hð  ·  ]! 
$  "  D 	$  "  / 	$  "   	$  "   		$ 4 5$  C4 5454455&  4 >  "  G    645444555&  s4	5$  4 >  "  G    	64	5$  4 >  "  G    	64 >  "  G    	64	5$  "  4544	55$  	 	4564	5$  "  4544	55$  	 	4	564	5$  "  4544	55$  	 	4	56CC       ¯      g  slot
	é g  t		Z g  t			W g  t		*	T g  t		<	Q g  top-card	rç  g  filenamef  osmosis.scm
 
	 		 		 		 		* 		* 		< 		< 		N 		^ 		_ 		i 		l 		r 		u 		| 		 	"  	  	  	
  	 ¦ 	+ © 	 ª 	 ± 	 ´ 	' · 	0 ¾ 	' ¿ 	" Á 	 Å 	 Æ 	 Ð 	 Ñ 	 ë 	7 ï 	 ð 	 ú 	 û 	 	7 	  	4 ¡	78 ¡	9 ¢	#C ¢	I £	#P ¤	#S ¤	2Z ¤	-\ ¤	#] £	a ¢	h ¥	/q ¥	r ¦	#| ¦	 §	# ¨	# ¨	2 ¨	- ¨	# §	 ¢	¡ ©	/« ©	¬ ª	#¶ ª	¼ «	#Ã ¬	#Æ ¬	2Í ¬	-Ï ¬	#Ð «	Ô ¢	Û ­	/å ­	 O	é  g  nameg  button-double-clicked CFR6;=2G>?<H 
       hÀ   è  ]		$  ­45$  045444555$  4 5"  "  Z454455$  A$  "  44	5455$  4 5"  "  $  C	 	6C      à      g  	from-slot
	 º g  card	 º g  slot-id		 º g  t		b  g  t	   ¸  g  filenamef  osmosis.scm
 ±
	 ²		 ²		 ³		 ³		 ´		 µ		! µ	'	$ µ	0	+ µ	'	, µ	"	. µ		/ ´		3 ´		4 ¶		G ·		N ·	'	Q ·	1	Y ·	'	Z ·		^ ·		b ¸		b ¸		p ¹		s ¹	0	z ¹	;	| ¹	0	} º	0  ¹	  ·	  »	   ³	 ¶ ¼	& ¸ ¼	 #	 º	  g  nameg  
placeable? CHR6H?I h@   û   ]	 (  C4 5$  "  4 4 55$  C 6  ó       g  id-list
		> g  t	,	>  g  filenamef  osmosis.scm
 ¾
	 ¿		 À		 À	"	 À		 À		 Á		! Á		" Á	*	' Á	8	) Á	*	, Á		, À		< Â		> Â	 		>  g  nameg  get-valid-move CIR"56IJ  h`   ó   ]4>   "  G    $  "  	 $  4		5"    $   C4	5  $   C6  ë       g  t
		) g  t
	:	^ g  t
	M	^  g  filenamef  osmosis.scm
 Ä
	 Å		 Æ		& Ç		- Æ		. È		5 È		: Æ		F É		M É		M Æ		\ Ê		^ Ê	 		^
  g  nameg  game-continuable CKR12  hP   þ   ] 	4455$  :	44	55$  %	44	55$  	44	55CCCCö       g  filenamef  osmosis.scm
 Ì
	 Í		 Í		 Í		 Í		 Í		 Î		 Î		" Î		# Î		' Í		* Ï		- Ï		5 Ï		6 Ï		: Í		= Ð		@ Ð		H Ð		I Ð	 		P
  g  nameg  game-won CLRIJ.MN       h8   Ø   ]45  $   C$  
45 C
45 C      Ð       g  t
			2  g  filenamef  osmosis.scm
 Ò
	 Ó		 Ó			 Ó			 Ó		 Ô		 Õ		  Õ		" Õ		% Õ		( Ö		, Ö		. Ö		1 Ö	 		2
  g  nameg  get-hint COR.P   h      ] 45  C      {       g  filenamef  osmosis.scm
 Û
	 Ü		 Ü			 Ü		 Ü		 Ü	 		
  g  nameg  get-options CQR    h      ]  C     x       g  options
		  g  filenamef  osmosis.scm
 Þ
	 ß			 ß	 		  g  nameg  apply-options CRR   h   R   ] C    J       g  filenamef  osmosis.scm
 â
 		
  g  nameg  timeout CSR4TiUi>  "  G  Vi$i7iAiCiFiKiLiOiQiRiSi@i6             g  filenamef  osmosis.scm		
		
¬	
i	8
(	=
	B
ÿ	F
	¬	M

ç	V
	]
	b
z	s
Y	w
Ó	{
¢ 
s ±
Æ ¾
6 Ä
 Ì
 Ë Ò
!~ Û
"# Þ
" â
" ä
"Ù æ
 	"Ù
   C6 