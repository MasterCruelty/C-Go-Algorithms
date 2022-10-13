
struct bi_node{
    int info;
    struct bi_node *next;
    struct bi_node *prev;
};

typedef struct bi_node *Bi_node;

typedef struct bi_list {
    Bi_node first;
    Bi_node last;
} *Bi_list;

int main(){
    struct bi_list l;
    return 0;
}
