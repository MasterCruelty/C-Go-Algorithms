Appunti dalle slide sugli alberi

inorder simmetrico: prima sottoalbero sx, poi la radice, poi sottoalbero dx
void bit_inorder ( Bit_node p ){
  if ( p ) {
    bit_inorder ( p -> l );
    bit_printnode ( p );
    bit_inorder ( p -> r );
  }
}

preorder anticipato: prima la radice, poi sottoalbero sx e poi sottoalbero dx
void bit_preorder ( Bit_node p ){
  if ( p ) {
    bit_printnode ( p );
    bit_preorder ( p -> l );
    bit_preorder ( p -> r );
  }
}


postoder differito: prima sottoalbero sx, poi sottoalbero dx e infine la radice
void bit_postorder ( Bit_node p ){
  if ( p ) {
    bit_postorder ( p -> l );
    bit_postorder ( p -> r );
    bit_printnode ( p );
  }
}


//conversione da array ad albero
//la radice è contenuta in a[0]
//il figlio sinistro è a[2*i+1]
//il figlio destro è a[2*i+2]
Bit_node bit_arr2tree(Item a[], int size,int j){
    if(j >= size|| a[j] == -1){
    	return NULL;
    }
    Bit_node root = bit_new(a[j]);
    root -> l = bit_arr2tree(a,size,(j+j+1));
    root -> r = bit_arr2tree(a,size,(j+j+2));
    return root;
}
