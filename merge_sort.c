



#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <malloc.h>

void heBin(int *a,int lenA,int *b,int lenB,int *res)
{
    int i=0;
    int j=0;
    while(i<lenA || j <lenB)
    {
        if(*a<*b)
        {
            *res = *a;
            res++;
            a++;
            i++;
        }else
        {
            *res = *b;
            res++;
            b++;
            j++;
        }

        if(i==lenA)
        {
            while(j<lenB)
            {
                *res = *b;
                res++;
                b++;
                j++;
            }
//            break;//

        }else if(j==lenB)
        {
            while(i<lenA)
            {
                *res = *a;
                res++;
                a++;
                i++;
            }
//            break;//
        }
    }
    return ;
}

void guiBin(int * intP,int len,int *res){
    int q = len / 2;
    int *arrA = NULL;
    int *arrB = NULL;
    if(q<1)
    {
        *res = *intP;
        return ;
    }

    arrA = (int*)malloc(sizeof(int)*(q));
    arrB = (int*)malloc(sizeof(int)*(len-q));
    
    //这里容易算错，不小心就溢出或者空间过大
    //0 1 2 3* 4 5 6 7 8
    guiBin(intP,q,arrA);//*0 , 4
    guiBin(intP+q,len-q,arrB);//*(0+4),9-4

    heBin(arrA,q,arrB,len-q,res);

    free(arrA);
    free(arrB);
    return ;

}

void guiBinStart(int * intP,int len)
{
    int *res = (int*)malloc(sizeof(int)*len);
    guiBin(intP,len,res);
    memcpy(intP,res,sizeof(int)*len);
    free(res);
    return ;
}

int main(){
    int arr[] ={1,2,9,6,8,5,3,4,5,9,10,546,65,234,45,123,654,2423,345};
    int i=0;
    int len=sizeof(arr) / sizeof(int);

    printf("first:\n");
    i=0;
    while(i<len)
    {

        printf("%d ",arr[i]);
        i++;
    }
    printf("\nlast:\n");


    guiBinStart(&arr[0],len);
    i=0;
    while(i<len)
    {

        printf("%d ",arr[i]);
        i++;
    }

    printf("\n");
    return 0;
}
