GOOF----LE-4-2.0£      ] . 4        h[      ] g  guile¤	 ¤	g  define-module*¤	 ¤	 ¤	g  ice-9¤	g  streams¤	 ¤		g  filenameS¤	
f  ice-9/streams.scm¤	g  exportsS¤	g  make-stream¤	g  
stream-car¤	g  
stream-cdr¤	g  stream-null?¤	g  list->stream¤	g  vector->stream¤	g  port->stream¤	g  stream->list¤	g  stream->reversed-list¤	g  stream->list&length¤	g  stream->reversed-list&length¤	g  stream->vector¤	g  stream-fold¤	g  stream-for-each¤	g  
stream-map¤	 ¤	g  set-current-module¤	 ¤	 ¤	g  make-promise¤	  ¤	! ¤	"g  force¤	#g  vector-length¤	$g  reverse!¤	%g  make-vector¤	&g  stream-fold-one¤	'g  stream-fold-many¤	(g  or-map¤	)g  map¤	*g  stream-for-each-one¤	+g  stream-for-each-many¤	,g  apply¤	-g  eof-object?¤C 5hø  ù   ]4	
5 4 >  "  G   !       h(   ¤   ]4LL 5  $   4L 5CC            g  o
			#  g  filenamef  ice-9/streams.scm
	V			W				W			X	
		X			Y			Z			Z			Z		 	Y	
	"	[	
 		#
   C h      ] O 6 ~       g  m
		 g  state		  g  filenamef  ice-9/streams.scm
	U
		V	 			  g  nameg  make-stream CR" h   ×   ]4 5C     Ï       g  stream
		  g  filenamef  ice-9/streams.scm
	]
		_		
	_	 		  g  nameg  
stream-carg  documentationf  BReturns the first element in STREAM.  This is equivalent to `car'. CR"        h   Ü   ]4 5C     Ô       g  stream
		  g  filenamef  ice-9/streams.scm
	a
		c		
	c	 		  g  nameg  
stream-cdrg  documentationf  GReturns the first tail of STREAM. Equivalent to `(force (cdr STREAM))'. CR"   h   ?  ]4 5C     7      g  stream
		  g  filenamef  ice-9/streams.scm
	e
		i			
	i	 		  g  nameg  stream-null?g  documentationf  ¨Returns `#t' if STREAM is the end-of-stream marker; otherwise
returns `#f'.  This is equivalent to `null?', but should be used
whenever testing for the end of a stream. CR h   T   ] C   L       g  l
		  g  filenamef  ice-9/streams.scm
	o	 		   C h   ù   ] 6      ñ       g  l
		
  g  filenamef  ice-9/streams.scm
	k
	
	n	 		
  g  nameg  list->streamg  documentationf  oReturns a newly allocated stream whose elements are the elements of
LIST.  Equivalent to `(apply stream LIST)'. CR#     h       ]	 L $  CL £ C           g  i
		 g  t		  g  filenamef  ice-9/streams.scm
	u			v			v			w			w	"		w	 		   C   h       ]	4 5 O 
6           g  v
		 g  len		  g  filenamef  ice-9/streams.scm
	r
		t			t			s	 		  g  nameg  vector->stream CR     hH     ]"  /45$  D4545"ÿÿÑ 
"ÿÿÃ     	      g  stream
		C g  s		5 g  acc			5 g  len			5  g  filenamef  ice-9/streams.scm
	z
		{			|			|			}			~			~	#	(	~		+	~	7	5	~		5	{		8	{		C	{	 		C  g  nameg  stream->reversed-list&length CR      h   ¬   ]4 >  G C ¤       g  stream
		 g  l		 g  len			  g  filenamef  ice-9/streams.scm
 
	 		 	 		  g  nameg  stream->reversed-list CR$ h    ¼   ]4 >  G 45D ´       g  stream
		 g  l		 g  len			  g  filenamef  ice-9/streams.scm
 
	 		 		 		 	 		  g  nameg  stream->list&length CR$ h   $  ]4 56         g  stream
		  g  filenamef  ice-9/streams.scm
 
	 		 	 		  g  nameg  stream->listg  documentationf  Returns a newly allocated list whose elements are the elements of STREAM.
If STREAM has infinite length this procedure will not terminate. CR% hX   5  ])4 >  G 45"  $(  "  %¤"ÿÿÜ
"ÿÿÑC     -      g  stream
		S g  l		S g  len			S g  v			S g  i		!	E g  l		!	E  g  filenamef  ice-9/streams.scm
 
	 		 		 		 		! 		' 			2 		6 	*	7 		: 		= 		E 		E 	 		S  g  nameg  stream->vector CR&'        h(   Ë   - 1 3 (  
 6 6  Ã       g  f
			& g  init			& g  stream				& g  rest				&  g  filenamef  ice-9/streams.scm
 
	 		 		$ 		& 	 			&	
	  g  nameg  stream-fold CR&      h0   Ò   ]45$  C 4 455456     Ê       g  f
		+ g  r		+ g  stream			+  g  filenamef  ice-9/streams.scm
  
	 ¡		 ¡		 £		 £		" £		# £	3	+ £	 			+	  g  nameg  stream-fold-one C&R('  h    ©   ] (  L C 4L  5C     ¡       g  cars
		  g  filenamef  ice-9/streams.scm
 ©	!	 «	#	 ¬	'	 ­	-	 ®	-	 ®	4	 ®	-	 ­	' 			  g  nameg  recur C)      hH   ù   ]45$  C 4 O Q 4455?456 ñ       g  f
		G g  r		G g  streams			G g  recur		"	:  g  filenamef  ice-9/streams.scm
 ¥
	 ¦		 ¦		 ©		" ©	!	- ª	.	7 ©	!	< ©		= ¯		G ¨	 		G	  g  nameg  stream-fold-many C'R*+    h(   ½   - 1 3 (   6 6      µ       g  f
			" g  stream			" g  rest				"  g  filenamef  ice-9/streams.scm
 ±
	 ²		 ³		  ´		" ´	 			"	
	  g  nameg  stream-for-each CR*    h8   Ç   ]45$  C4 45>  "  G   456     ¿       g  f
		3 g  stream		3  g  filenamef  ice-9/streams.scm
 ¶
	 ·		 ·		 ¹		 ¹		 ¹		+ º		3 º	 			3	  g  nameg  stream-for-each-one C*R(,)+    h@   É   ]45$  C4 45>  "  G   456     Á       g  f
		; g  streams		;  g  filenamef  ice-9/streams.scm
 ¼
	 ½		 ½		 ¿		 ¿		$ ¿		1 À	 	; À	 			;	  g  nameg  stream-for-each-many C+R       h0   ¢   ]	4 5$  C4L 4 554 5C              g  s
		) g  t			)  g  filenamef  ice-9/streams.scm
 Ç		 È			 È		 É		 É	"	  É		! É	2	( É	 			)   C()        h0   ¨   ]	4 5$  C4L 4 5?4 5C         g  streams
		/ g  t		/  g  filenamef  ice-9/streams.scm
 Ë		 Ì		 Ì		 Í		 Í	(	$ Í		% Î		. Í	 			/   C     h0   \  - 1 3 (   O 6 O 6    T      g  f
			, g  stream			, g  rest				,  g  filenamef  ice-9/streams.scm
 Â
	 Æ		 Ç		* Ï		, Ë	 			,	
	  g  nameg  
stream-mapg  documentationf  Returns a newly allocated stream, each element being the result of
invoking F with the corresponding elements of the STREAMs
as its arguments. CR-        h(      ]4L  545$  C C            g  p
		# g  o			# g  t			#  g  filenamef  ice-9/streams.scm
 Ò		 Ó			 Ó		 Ô		 Ô		" Õ	 		#   C      h      ]O  6        g  port
		 g  read		  g  filenamef  ice-9/streams.scm
 Ñ
	 Ò	 			  g  nameg  port->stream CRCñ       g  m
		(  g  filenamef  ice-9/streams.scm		
±	U
ª	]
¯	a
	e
	k
	r
		z

g 
W 
 
@ 
N 
m  
¤ ¥
  ±
º ¶
ä ¼
g Â
ö Ñ
 	ø
   C6 