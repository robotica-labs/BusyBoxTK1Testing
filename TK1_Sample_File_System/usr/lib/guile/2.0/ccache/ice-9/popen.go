GOOF----LE-4-2.08      ] > 4        hn      ] g  guile�	 �	g  define-module*�	 �	 �	g  ice-9�	g  popen�	 �		g  filenameS�	
f  ice-9/popen.scm�	g  exportsS�	g  port/pid-table�	g  
open-pipe*�	g  	open-pipe�	g  
close-pipe�	g  open-input-pipe�	g  open-output-pipe�	g  open-input-output-pipe�	 �	g  set-current-module�	 �	 �	g  load-extension�	g  string-append�	f  	libguile-�	g  effective-version�	f  scm_init_popen�	g  make-soft-port�	g  
write-char�	g  display�	g  force-output�	 g  	read-char�	!g  
close-port�	"f  r+�	#g  make-rw-port�	$g  make-guardian�	%g  pipe-guardian�	&g  make-weak-key-hash-table�	'g  apply�	(g  open-process�	)g  %make-void-port�	*g  
hashq-set!�	+f  /bin/sh�	,f  -c�	-g  	hashq-ref�	.g  hashq-remove!�	/g  	fetch-pid�	0g  waitpid�	1g  close-process�	2g  catch�	3g  system-error�	4g  WNOHANG�	5g  close-process-quietly�	6g  error�	7f  close-pipe: pipe not in table�	8g  
reap-pipes�	9g  	add-hook!�	:g  after-gc-hook�	;g  	OPEN_READ�	<g  
OPEN_WRITE�	=g  	OPEN_BOTH�C 5     hP    ]4	
5 4 >  "  G   4i4i4i5 5>  "  G     h   Z   ] L 6      R       g  c
		
  g  filenamef  ice-9/popen.scm�
		��	
		�� 		
   C        h   Z   ] L 6      R       g  s
		
  g  filenamef  ice-9/popen.scm�
		��	
		�� 		
   C        h   J   ] L 6B       g  filenamef  ice-9/popen.scm�
	 	��		 	�� 		
   C         h   J   ] L 6B       g  filenamef  ice-9/popen.scm�
	!	��		!	�� 		
   C!        h    R   ] 4L>  "  G  L 6      J       g  filenamef  ice-9/popen.scm�
	"	��		"	��		"	&�� 		
   C" h0   �   ]O O O  O  O  6�       g  	read-port
		0 g  
write-port		0  g  filenamef  ice-9/popen.scm�
	
��	,		��	.	#	��	0		�� 		0	  g  nameg  make-rw-port� C#R4$i5 %R4&i	5R'(#)%*   h�   �  - 1 3 4 >  G $  $  45"  "  $  "  $  "  $  "  4 54>  "  G  4>  "  G  C �      g  mode
		 � g  command		 � g  args			 � g  	read-port		 � g  
write-port		 � g  pid		 � g  t		B	r g  port		r �  g  filenamef  ice-9/popen.scm�
	,
��	
	4	��		3	��	)	6	��	0	7	��	B	6	��	i	:	��	r	6	��	u	;	�� �	<	�� 		 �	
	  g  nameg  
open-pipe*�g  documentationf [Executes the program @var{command} with optional arguments
@var{args} (all strings) in a subprocess.
A port to the process (based on pipes) is created and returned.
@var{mode} specifies whether an input, an output or an input-output
port to the process is created: it should be the value of
@code{OPEN_READ}, @code{OPEN_WRITE} or @code{OPEN_BOTH}.� CR+,   h   �  ] 6  �      g  command
		 g  mode		  g  filenamef  ice-9/popen.scm�
	?
��		E	��	
	E	��		E	�� 			  g  nameg  	open-pipe�g  documentationf ;Executes the shell command @var{command} (a string) in a subprocess.
A port to the process (based on pipes) is created and returned.
@var{mode} specifies whether an input, an output or an input-output
port to the process is created: it should be the value of
@code{OPEN_READ}, @code{OPEN_WRITE} or @code{OPEN_BOTH}.� CR-.   h(   �   ]	4 54 >  "  G  C    �       g  port
		$ g  pid		$  g  filenamef  ice-9/popen.scm�
	G
��		H	��		H	��		I	�� 		$  g  nameg  	fetch-pid� C/R!0   h    �   ]4 �>  "  G  4 �5�C �       g  port/pid
		  g  filenamef  ice-9/popen.scm�
	L
��		M	��		M	��		M	��		N	��		N	��		N	��		N	�� 			  g  nameg  close-process� C1R23!        h   R   ] L �6       J       g  filenamef  ice-9/popen.scm�
	V		��		W	��			W	�� 			
   C    h   W   -  1  3 C     O       g  args
			  g  filenamef  ice-9/popen.scm�
	X		�� 			


   C04%*   h8   �   ]4L �5  �
�$  4L �>  "  G  L �L �6C�       g  
pid/status
		8  g  filenamef  ice-9/popen.scm�
	Z		��		[	��		[	&��		[	��		[	��		\	��		\	��		\	��		^	��		^	#��	"	^	��	1	`	 ��	4	`	/��	6	_	�� 		8
   C   h   W   -  1  3 C     O       g  args
			  g  filenamef  ice-9/popen.scm�
	a		�� 			


   C      h0   �   ]4 O >  "  G   O 6    �       g  port/pid
		,  g  filenamef  ice-9/popen.scm�
	T
��		U	��		U		��		U	��	!	Y		��	,	Y	�� 		,  g  nameg  close-process-quietly� C5R/671       h0   �  ]	4 5$  "  4>  "  G   �6�      g  p
		0 g  pid			0  g  filenamef  ice-9/popen.scm�
	c
��		g	��			g	��		h	��		i	��		i	��		i	��	.	j	��	0	j	�� 
		0  g  nameg  
close-pipe�g  documentationf  �Closes the pipe created by @code{open-pipe}, then waits for the process
to terminate and returns its status value, @xref{Processes, waitpid}, for
information on how to interpret this value.� CR/5%        hP   �   ]"  > $  64 5$  4 �>  "  G  "   45  "���C45  "��� �       g  p
		D g  pid		7  g  filenamef  ice-9/popen.scm�
	m	��		n	��		o	��		q	��		q	��		r	��		s	��	#	s	*��	(	s	��	8	t	��	B	t	��	D	n	��	E	n	��	O	n	�� 		O
  g  nameg  
reap-pipes� C8R49i:i8i>  "  G  ;  h   �   ] 6      �       g  command
		
  g  filenamef  ice-9/popen.scm�
	x
��	
	z	�� 		
  g  nameg  open-input-pipe�g  documentationf  9Equivalent to @code{open-pipe} with mode @code{OPEN_READ}� CR<   h   �   ] 6      �       g  command
		
  g  filenamef  ice-9/popen.scm�
	|
��	
	~	�� 		
  g  nameg  open-output-pipe�g  documentationf  :Equivalent to @code{open-pipe} with mode @code{OPEN_WRITE}� CR= h   �   ] 6      �       g  command
		
  g  filenamef  ice-9/popen.scm�
 �
��	
 �	�� 		
  g  nameg  open-input-output-pipe�g  documentationf  9Equivalent to @code{open-pipe} with mode @code{OPEN_BOTH}� CRC      g  m
		(  g  filenamef  ice-9/popen.scm�		
��	)		��	.		��	4		!��	5		-��	=		��	?		��	D		���	
���	'	���	'
���	*	���	*
��	,
��	#	?
��	�	G
��
�	L
��5	T
��	c
��\	l
��]	v
��e	x
��W	|
��N �
�� 	P
   C6 