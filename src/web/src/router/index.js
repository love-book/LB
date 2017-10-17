import Vue from 'vue'
import Router from 'vue-router'
import tabber from '@/components/tabber'
import navbar from '@/components/navbar'
import near from '@/components/near'
import bookstore from '@/components/bookstore'
import bookcase from '@/components/bookcase'
import mine from '@/components/mine'
import exchange from '@/components/exchange/exchange'
import book from '@/components/book'
import bookinfo from '@/components/bookinfo'
import message from '@/components/message'
import collectbookcase from '@/components/collectbookcase'
import shelvesrecord from '@/components/shelvesrecord'
import collectbooks from '@/components/collectbooks'
import editprofile from '@/components/editprofile'
import addcontact from '@/components/addcontact'
import editbookrs from '@/components/editbookrs'
import bindphonenum from '@/components/bindphonenum'
import getcontact from '@/components/getcontact'
import signature from '@/components/signature'
import ideasbox from '@/components/ideasbox'
import choosesex from '@/components/choosesex'
import aboutus from '@/components/aboutus'
import otherbookcase from '@/components/otherbookcase'
import contactway from '@/components/contactway'
import choosebook from '@/components/choosebook'
import delbook from '@/components/delbook'
import editbooktp from '@/components/editbooktp'
Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'tabber',
      redirect: '/tabber/navbar/near',
      component: tabber,
      children:[
                { 
                  path: '/tabber/navbar', 
                  redirect: '/tabber/navbar/near',
                  component: navbar,
                  children:[
                    { 
                      path: '/tabber/navbar/near', 
                      component: near,
                    },
                    { 
                      path: '/tabber/navbar/bookstore', 
                      component: bookstore,
                    },
                  ]
                },
                { 
                  path: '/tabber/container2', 
                  component: bookcase,
                  children:[]
                },
                { 
                  path: '/tabber/container3', 
                  component: mine,
                  children:[]
                },
            ]
    },
    {
      path: '/book/:bookid',
      name: 'book',
      component: book,
    },
    { 
        path: '/otherbookcase/:userid', 
        name: 'otherbookcase',
        component: otherbookcase,
    },
    { 
        path: '/bookinfo/:bookqid', 
        name: 'bookinfo',
        component: bookinfo,
    },
    { 
        path: '/getcontact/:bookqid', 
        name: 'getcontact',
        component: getcontact,
    },
    { 
        path: '/contactway/:userid', 
        name: 'contactway',
        component: contactway,
    },
    { 
        path: '/choosebook/:userid', 
        name: 'choosebook',
        component: choosebook,
    },
    { 
        path: '/signature', 
        name: 'signature',
        component: signature,
    },
    { 
        path: '/bindphonenum', 
        name: 'bindphonenum',
        component: bindphonenum,
    },
    { 
        path: '/delbook', 
        name: 'delbook',
        component: delbook,
    },
    { 
        path: '/editbookrs/:isbn', 
        name: 'editbookrs',
        component: editbookrs,
    },
    { 
        path: '/editbooktp', 
        name: 'editbooktp',
        component: editbooktp,
    },
    // {
    //   path: '/:userid',
    //   name: 'gootherbookcase',
    //   component: otherbookcase,
    // },
    // { 
    //     path: '/:userid/:bookqid', 
    //     name: 'pgobookinfo',
    //     component: bookinfo,
    // },
    {
      path: '/editprofile',
      name: 'editprofile',
      component: editprofile
    },
    {
      path: '/collectbooks',
      name: 'collectbooks',
      component: collectbooks
    },
    {
      path: '/collectbookcase',
      name: 'collectbookcase',
      component: collectbookcase
    },
    {
      path: '/message',
      name: 'message',
      component: message
    },
    {
    	path: '/exchange',
    	name: 'exchange',
    	component:exchange
    },
    {
    	path:'/shelvesrecord',
    	name:'shelvesrecord',
    	component:shelvesrecord
    },
    {
    	path:'/ideasbox',
    	name:'ideasbox',
    	component:ideasbox
    },
    {
    	path:'/aboutus',
    	name:'aboutus',
    	component:aboutus
    },
    {
      path:'*',
      redirect: '/',
    }
  ],
  linkActiveClass: 'active'
})
