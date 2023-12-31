GOOF----LE-4-2.0Î      ] W 4       hL      ] g  guile¤	 ¤	g  define-module*¤	 ¤	 ¤	g  language¤	g  elisp¤	g  parser¤		 ¤	
g  filenameS¤	f  language/elisp/parser.scm¤	g  importsS¤	g  lexer¤	 ¤	 ¤	 ¤	g  exportsS¤	g  
read-elisp¤	 ¤	g  set-current-module¤	 ¤	 ¤	g  error¤	g  parse-error¤	g  
make-fluid¤	g  circular-definitions¤	g  make-hash-table¤	g  make-circular-definitions¤	g  circular-ref¤	f  invalid token for circular-ref¤	g  	hashq-ref¤	 f  undefined circular reference¤	!g  circular-def¤	"f  "invalid token for circular-define!¤	#g  
hashq-set!¤	$g  make-promise¤	%$ ¤	&$ ¤	'g  circular-define!¤	(g  promise?¤	)g  force¤	*g  force-promises!¤	+g  vector-length¤	,g  finish¤	-f  'lexer-buffer is not empty when finished¤	.g  peek¤	/g  get¤	0f  invalid lexer-buffer action¤	1g  make-lexer-buffer¤	2g  square-close¤	3g  paren-close¤	4f  got different token than peeked¤	5g  get-expression¤	6g  get-list¤	7g  dot¤	8g  length¤	9f  &expected exactly one element after dot¤	:g  quote¤	;::¤	<g  	backquote¤	=g  `¤	><=¤	?g  unquote¤	@g  ,¤	A?@¤	Bg  unquote-splicing¤	Cg  ,@¤	DBC¤	E;>AD ¤	Fg  quotation-symbols¤	Gg  eof¤	Hf  end of file during parsing¤	Ig  integer¤	Jg  float¤	Kg  symbol¤	Lg  	character¤	Mg  string¤	Ng  set-source-properties!¤	Og  source-properties¤	Pg  function¤	Qg  assq-ref¤	Rg  
paren-open¤	Sg  square-open¤	Tg  list->vector¤	Uf  expected expression, got¤	Vg  	get-lexer¤C 5       h0  Ì   ]4	
5 4 >  "  G         h   ¥   - 1 3 @              g  token
			 g  msg			 g  args				  g  filenamef  language/elisp/parser.scm
	"
		#	 				
	  g  nameg  parse-error CR4i5 R        h   y   ] 6   q       g  filenamef  language/elisp/parser.scm
	3
		4	 		
  g  nameg  make-circular-definitions CR   hH     ] &  "  4 >  "  G   4[5$  C 6 ú       g  token
		G g  id	&	G g  value		2	G  g  filenamef  language/elisp/parser.scm
	6
		7			7			7			8			8			8		&	9		&	9		)	:		2	9		:	;		C	=		G	=	 		G  g  nameg  circular-ref CR!"#&    h   L   ] M C   D       g  filenamef  language/elisp/parser.scm
	I	 		
   C#      h   u   ] NL L 6m       g  
real-value
		  g  filenamef  language/elisp/parser.scm
	J			K			L	 		   C 	       h`   	  ] &  "  4 >  "  G  [ H44O 5>  "  G  O C          g  token
		\ g  value	*	\ g  table		*	\ g  id		*	\  g  filenamef  language/elisp/parser.scm
	C
		D			D			D			E			E			E		*	H		*	F		1	I		8	I		H	I	 		\  g  nameg  circular-define! C'R()*+        hÀ   õ  ] $  N4 5$   4 5"  4 >  "  G  4 5$   4 5C 6 $  \4 5"  H$  = £45$   45¤"  4>  "  G  "ÿÿºC
"ÿÿ±C    í      g  data
	 ¼ g  len	e º g  i		k ³ g  el		y ±  g  filenamef  language/elisp/parser.scm
	U
		W				V		
	Y	
		Y			Y	
		Y			Z			Z	 		Z		 	Z	
	%	[	
	*	[		/	[	
	8	\	
	=	\		?	\	
	C	\		F	]		K	]	 	M	]		N	]	
	U	^		W	^	
	Z	_		^	V		_	`		e	`		k	a		p	b		t	b		y	c		y	c		|	d	 	d	 	e	& 	e	 	f	 «	g	 ±	g	 ³	a	 *	 ¼  g  nameg  force-promises! C*R,-./0       hP   ß   ]	 &  M$  6CM$  "  4L 5 N $  MC $  
MNC 6×       g  action
		P g  result	@	H  g  filenamef  language/elisp/parser.scm
	t			u		
	u			v	
		w			w			z		#	{	!	)	{		2	|		@			E 		L 		P 	 		P   C      h      ]	H O C        g  lex
		 g  
look-ahead		  g  filenamef  language/elisp/parser.scm
	r
		s	 		  g  nameg  make-lexer-buffer C1R.23/456789       hà      ]#4 5$  "  &  &4 5&  "  4>  "  G  C"  4 54 5C$  e	&  Y4 5&  "  4>  "  G  4 54
5$  "  4>  "  G  C"ÿÿy"ÿÿu             g  lex
	 Ù g  	allow-dot	 Ù g  close-square		 Ù g  next			 Ù g  type		 Ù g  head		U	j g  tail		b	j g  tail	 ¦ Ñ  g  filenamef  language/elisp/parser.scm
 
	 		 			 			 		 		 		 		 	!	  	/	$ 		% 		) 	 	+ 		, 		2 		7 		; 		@ 		I 		O 		U 		X  		b 		i ¡		j 		t 		x 		y 		} 	 	 	  	  	  	  	  	  	 ¦ 	 © 	 ± 	 µ 	
 º 	 À 	 Å 	 Ð 	
 .	 Ù	  g  nameg  get-list C6REFR/GHIJKLMNOP5:<?BQFR6ST!'*U  h  »  ]!4 5$  6$  "  /$  "  !$  "  $  "  	$  ,$  4
45>  "  G  "   C$  54 5 $  4
45>  "  G  "   C$  "  !$  "  $  "  $  <454 5 $  4
45>  "  G  "   C$  24 5$  4
45>  "  G  "   C$  744 55$  4
45>  "  G  "   C$  6$  9454 54>  "  G  4>  "  G  C6       ³      g  lex
	 g  token		 g  type		 g  result		e  g  result	 £ Ì g  result	; g  result	Mv g  result	¶ g  setter	Õ g  expr	Þ  
g  filenamef  language/elisp/parser.scm
 «
	 ¬		 ¬			 ¬			 ¬		 ­		 ¬		 ´		 ¶		! ¶		* ´		e ¸		e ¸		j ¯		n ¯		o °		t ²		 °	  ´	  º	  º	 £ º	 £ º	 ¨ ¯	 ¬ ¯	 ­ °	 ² ²	 ½ °	 Õ ´	  ¼		 ½	 ¼	 ¼	 ¯	 ¯	 °	! ²	, °	D ´	E ¿	M ¿	R ¯	V ¯	W °	\ ²	g °	 ´	 Á	 Á	 Á	 Á	 ¯	 ¯	 °	 ²	§ °	¿ ´	Å Ã	Î ´	Ï Æ	Õ Æ	Ø Ç	Þ Æ	á È		ó É		 Ì	 Ì	 D	  g  nameg  get-expression C5RV1.G5, 	     hX   d  ]!45 Y4 54545&  "  454>  "  G  ZCZF \      g  port
		W g  lexer		S g  lexbuf			S g  next		$	S g  result		=	S  g  filenamef  language/elisp/parser.scm
 Ò
	 Ó	&	 Ô		 Ô		 Õ		 Ô		 Ö		" Ö		$ Ö		$ Ô		) ×		+ ×		/ ×		2 Ø	
	7 Ù		= Ù	
	@ Ú		D Ú		I Ú	 		W  g  nameg  
read-elisp CRC       Ä       g  m
		,  g  filenamef  language/elisp/parser.scm		
 	"
	1	
	1
¤	3
	6
	C
h	U

u	r
# 
% ¦	( ¦
F «
' Ò
 	)
   C6 